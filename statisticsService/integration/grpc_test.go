package integration

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/Denterry/SocialNetwork/statisticsService/pkg/stat_v1"
	"github.com/docker/compose/v2/pkg/api"
	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Test for checking gRPC routes of statistics service
func TestWithStatsGRPC(t *testing.T) {
	identifier := tc.StackIdentifier("some_new_ident")
	compose, err := tc.NewDockerComposeWith(tc.WithStackFiles("../../docker-compose.kafka.yml", "../../docker-compose.cheap.yml"), identifier)
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
		} else if err.Error() == "container stats-db-migration exited (0)" {
			log.Println("Container stats-db-migration exited successfully")
		} else if err.Error() == "container auth-db-migration exited (0)" {
			log.Println("Container auth-db-migration exited successfully")
		} else {
			require.NoError(t, err, "compose.Up()")
		}
	}

	service, err := compose.ServiceContainer(ctx, "statistics-service")
	require.NoError(t, err, "compose.ServiceContainer()")

	time.Sleep(time.Second * 10)

	serviceHost, err := service.Host(ctx)
	require.NoError(t, err, "service.Host()")

	// endpoints, err := service.Endpoint(ctx, "")
	// require.NoError(t, err, "service.Endpoint()")

	// fmt.Println(endpoints)

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%s", serviceHost, "8083"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err, "grpc.NewClient()")

	client := stat_v1.NewStatisticsServiceClient(conn)
	_ = client

	time.Sleep(time.Second * 10)

	res, err := client.TotalViewsLikes(ctx, &stat_v1.TotalViewsLikesRequest{
		PostId: int64(1),
	})
	require.NoError(t, err, "client.TotalViewsLikes()")
	require.NotNil(t, res, "res")
	require.Equal(t, res.PostId, int64(1))
	require.Equal(t, uint64(0x0), res.TotalViews)
	require.Equal(t, uint64(0x0), res.TotalLikes)

	res, err = client.TotalViewsLikes(ctx, nil)
	require.Error(t, err, "client.TotalViewsLikes()")
	require.Nil(t, res, "res")
	require.Equal(t, err.Error(), "rpc error: code = InvalidArgument desc = Post Id must be specified")
}
