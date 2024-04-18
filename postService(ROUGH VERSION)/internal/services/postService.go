package services

import (
	"context"
	"log/slog"
	"time"

	"github.com/Denterry/SocialNetwork/postService/internal/domain/models"
	"github.com/Denterry/SocialNetwork/postService/internal/grpc/server"
)

type Postim struct {
	log      *slog.Logger
	storage  Storage
	tokenTTL time.Duration
}

type Storage interface {
	Create(
		ctx context.Context,
		authorId int,
		title string,
		content string,
	) (PostId int, err error)

	Update(
		ctx context.Context,
		postId int,
		authorId int,
		title string,
		content string,
	) (PostId int, err error)

	Get(
		ctx context.Context,
		postId int,
	) (AuthorId int, Title string, Content string, err error)

	Delete(
		ctx context.Context,
		postId int,
	) (PostId int, err error)

	// GetAll(

	// )
}

type AppProvider interface {
	App(ctx context.Context, appID int) (models.App, error)
}

// NewAppProvider creates a new instance of the AppProvider type using the provided database
func NewPostim(
	log *slog.Logger,
	storage Storage,
	tokenTTL time.Duration,
) *Postim {
	return &Postim{
		log:      log,
		storage:  storage,
		tokenTTL: tokenTTL,
	}
}

func CreatePost(
	ctx context.Context,
	authorId int64,
	title string,
	content string,
) (PostId int64, Title string, Content string, AuthorId string, err error) {
	const op = "services.postService.CreatePost"

	panic("not implemented")
}

func UpdatePost(
	ctx context.Context,
	postrId int64,
	title string,
	content string,
) (PostId int64, Title string, Content string, AuthorId string, err error) {
	const op = "services.postService.CreatePost"

	panic("not implemented")
}

func DeletePost(
	ctx context.Context,
	postrId int64,
) (PostId int64, Title string, Content string, AuthorId string, err error) {
	const op = "services.postService.CreatePost"

	panic("not implemented")
}

func GetPostById(
	ctx context.Context,
	postrId int64,
) (PostId int64, Title string, Content string, AuthorId string, err error) {
	const op = "services.postService.CreatePost"

	panic("not implemented")
}

func GetPosts(
	ctx context.Context,
	page int32,
	pageSize int32,
) (Posts []*server.PostResp, err error) {
	const op = "services.postService.CreatePost"

	panic("not emplemented")
}
