package integration

import (
	"context"
	"log"
	"testing"

	"github.com/Denterry/SocialNetwork/postService/pkg/post_v1"
	"github.com/docker/compose/v2/pkg/api"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	DB_HOST     = "post_database"
	DB_PORT     = "5432"
	DB_USER     = "postiniter"
	DB_PASSWORD = "qwerty123456"
	DB_NAME     = "post_db"
)

func TestWithPostsGRPC(t *testing.T) {

	// ctx := context.Background()
	// req := testcontainers.ContainerRequest{
	// 	FromDockerfile: testcontainers.FromDockerfile{
	// 		Context:    "../.",
	// 		Dockerfile: "post.Dockerfile",
	// 		BuildArgs: map[string]*string{
	// 			"DB_HOST":     &DB_HOST,
	// 			"DB_PORT":     &DB_PORT,
	// 			"DB_USER":     &DB_USER,
	// 			"DB_PASSWORD": &DB_PASSWORD,
	// 			"DB_NAME":     &DB_NAME,
	// 		},
	// 	},
	// 	ExposedPorts: []string{"8081/tcp"},
	// 	WaitingFor:   wait.ForLog("Ready to accept connections"),
	// }

	identifier := tc.StackIdentifier("some_ident")
	compose, err := tc.NewDockerComposeWith(tc.WithStackFiles("../../integration.post.docker-compose.yml"), identifier)
	require.NoError(t, err, "NewDockerComposeAPIWith()")
	_ = compose

	t.Cleanup(func() {
		require.NoError(t, compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal), "compose.Down()")
	})

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	err = compose.Up(ctx, tc.WithRecreate(api.RecreateNever), tc.Wait(true))
	if err != nil {
		if err.Error() == "container post-db-migration exited (0)" {
			log.Println("Container post-db-migration exited successfully")
		} else {
			require.NoError(t, err, "compose.Up()")
		}
	}

	service, err := compose.ServiceContainer(ctx, "post-service")
	require.NoError(t, err, "compose.ServiceContainer()")

	endpoints, err := service.Endpoint(ctx, "")
	require.NoError(t, err, "service.Endpoint()")

	conn, err := grpc.NewClient(endpoints, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err, "grpc.NewClient()")

	client := post_v1.NewPostServiceClient(conn)

	authorID := uuid.Nil
	authorIDStr := authorID.String()
	res, err := client.CreatePost(ctx, &post_v1.CreatePostRequest{
		Title:    "Test Title",
		Content:  "Test Content",
		AuthorId: authorIDStr,
	})
	require.NoError(t, err, "client.CreatePost()")
	require.NotNil(t, res.PostId, "res.PostId")
	require.Equal(t, res.Title, "Test Title")
	require.Equal(t, res.Content, "Test Content")
	require.Equal(t, res.AuthorId, authorIDStr)

	_, err = client.GetPostById(ctx, &post_v1.GetPostByIdRequest{
		PostId:   0,
		AuthorId: (uuid.Nil).String(),
	})
	require.Error(t, err, "client.GetPostById()")
	require.Equal(t, err.Error(), "rpc error: code = InvalidArgument desc = Post Id must be specified")

	_, err = client.GetPostById(ctx, &post_v1.GetPostByIdRequest{
		PostId:   10000,
		AuthorId: (uuid.Nil).String(),
	})
	require.Error(t, err, "client.GetPostById()")
	require.Equal(t, err.Error(), "rpc error: code = Unknown desc = sql: no rows in result set")
}
