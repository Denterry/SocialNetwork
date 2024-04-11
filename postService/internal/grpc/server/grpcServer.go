package server

import (
	"context"
	"fmt"

	"github.com/Denterry/SocialNetwork/postService/internal/domain/models"
	"github.com/Denterry/SocialNetwork/postService/pkg/post_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PostRepository provides an interface for CRUD operations on posts/tasks
type PostRepository interface {
	CreatePost(ctx context.Context, post models.Post) (*models.Post, error)
	UpdatePost(ctx context.Context, post models.Post) (*models.Post, error)
	DeletePost(ctx context.Context, id int64) error
	GetPostById(ctx context.Context, id int64) (*models.Post, error)
	ListPosts(ctx context.Context, page, pageSize int) ([]*models.Post, error)
}

type ServerAPI struct {
	post_v1.UnimplementedPostServiceServer
	repo PostRepository
}

func NewServerAPI(repo PostRepository) *ServerAPI {
	return &ServerAPI{repo: repo}
}

func (sapi *ServerAPI) CreatePost(ctx context.Context, request *post_v1.CreatePostRequest) (*post_v1.PostResponse, error) {
	if request.GetAuthorId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Author Id must be specified")
	}

	if request.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "Title must be specified")
	}

	if request.GetContent() == "" {
		return nil, status.Error(codes.InvalidArgument, "Content must be specified")
	}

	post := models.Post{
		ID:       -1,
		AuthorID: request.GetAuthorId(),
		Title:    request.GetTitle(),
		Content:  request.GetContent(),
	}

	res, err := sapi.repo.CreatePost(ctx, post)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	fmt.Print("The post was created successfully!")
	return &post_v1.PostResponse{
		PostId:   res.ID,
		Title:    res.Title,
		Content:  res.Content,
		AuthorId: res.AuthorID,
	}, nil
}

// func (sapi *ServerAPI) UpdatePost(ctx context.Context, request *post_v1.UpdatePostRequest) (*post_v1.PostResponse, error) {
// 	if request.GetPostId() == 0 {
// 		return nil, status.Error(codes.InvalidArgument, "Author Id must be specified")
// 	}

// 	if request.GetTitle() == "" {
// 		return nil, status.Error(codes.InvalidArgument, "Title must be specified")
// 	}

// 	if request.GetContent() == "" {
// 		return nil, status.Error(codes.InvalidArgument, "Content must be specified")
// 	}

// 	return &post_v1.PostResponse{}, nil
// }

// func (sapi *ServerAPI) DeletePost(ctx context.Context, request *post_v1.DeletePostRequest) (*post_v1.PostResponse, error) {
// 	if request.GetPostId() == 0 {
// 		return nil, status.Error(codes.InvalidArgument, "Author Id must be specified")
// 	}

// 	return &post_v1.PostResponse{}, nil
// }

// func (sapi *ServerAPI) GetPostById(ctx context.Context, request *post_v1.GetPostByIdRequest) (*post_v1.PostResponse, error) {
// 	if request.GetPostId() == 0 {
// 		return nil, status.Error(codes.InvalidArgument, "Author Id must be specified")
// 	}

// 	return &post_v1.PostResponse{}, nil
// }

// func (sapi *ServerAPI) GetPosts(ctx context.Context, request *post_v1.GetPostsRequest) (*post_v1.GetPostsResponse, error) {
// 	if request.GetPage() == 0 {
// 		return nil, status.Error(codes.InvalidArgument, "Page must be specified")
// 	}

// 	if request.GetPageSize() == 0 {
// 		return nil, status.Error(codes.InvalidArgument, "Page Size must be specified")
// 	}

// 	return &post_v1.GetPostsResponse{}, nil
// }
