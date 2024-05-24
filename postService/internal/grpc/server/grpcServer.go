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
	GetListPosts(ctx context.Context, page, pageSize int) ([]*models.Post, error)
	GetAuthorByPostId(ctx context.Context, postID int64) (int64, error)
}

type serverAPI struct {
	post_v1.UnimplementedPostServiceServer
	repo PostRepository
}

func NewServerAPI(repo PostRepository) *serverAPI {
	return &serverAPI{repo: repo}
}

func (sapi *serverAPI) CreatePost(ctx context.Context, request *post_v1.CreatePostRequest) (*post_v1.PostResponse, error) {
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

func (sapi *serverAPI) UpdatePost(ctx context.Context, request *post_v1.UpdatePostRequest) (*post_v1.PostResponse, error) {
	if request.GetPostId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Post Id must be specified")
	}

	if request.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "Title must be specified")
	}

	if request.GetContent() == "" {
		return nil, status.Error(codes.InvalidArgument, "Content must be specified")
	}

	correctAuthorId, err := sapi.repo.GetAuthorByPostId(ctx, request.GetPostId())
	if err != nil {
		return nil, err
	}

	if request.GetAuthorId() != correctAuthorId {
		return nil, fmt.Errorf("author id does not match the author of this post")
	}

	post := models.Post{
		ID:       request.GetPostId(),
		AuthorID: request.GetAuthorId(),
		Title:    request.GetTitle(),
		Content:  request.GetContent(),
	}

	res, err := sapi.repo.UpdatePost(ctx, post)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	fmt.Print("The post was updated successfully!")
	return &post_v1.PostResponse{
		PostId:   res.ID,
		Title:    res.Title,
		Content:  res.Content,
		AuthorId: res.AuthorID,
	}, nil
}

func (sapi *serverAPI) DeletePost(ctx context.Context, request *post_v1.DeletePostRequest) (*post_v1.PostResponse, error) {
	if request.GetPostId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Post Id must be specified")
	}

	correctAuthorId, err := sapi.repo.GetAuthorByPostId(ctx, request.GetPostId())
	if err != nil {
		return nil, err
	}

	if request.GetAuthorId() != correctAuthorId {
		return nil, fmt.Errorf("author id does not match the author of this post")
	}

	err = sapi.repo.DeletePost(ctx, request.GetPostId())
	if err != nil {
		return nil, err
	}

	fmt.Print("The post was deleted successfully!")
	return &post_v1.PostResponse{}, nil
}

func (sapi *serverAPI) GetPostById(ctx context.Context, request *post_v1.GetPostByIdRequest) (*post_v1.PostResponse, error) {
	if request.GetPostId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Post Id must be specified")
	}

	correctAuthorId, err := sapi.repo.GetAuthorByPostId(ctx, request.GetPostId())
	if err != nil {
		return nil, err
	}

	if request.GetAuthorId() != correctAuthorId {
		return nil, fmt.Errorf("author id does not match the author of this post")
	}

	res, err := sapi.repo.GetPostById(ctx, request.GetPostId())
	if err != nil {
		return nil, err
	}

	fmt.Print("The post was goten successfully!")
	return &post_v1.PostResponse{
		PostId:   res.ID,
		AuthorId: res.AuthorID,
		Title:    res.Title,
		Content:  res.Content,
	}, nil
}

func (sapi *serverAPI) GetListPosts(ctx context.Context, request *post_v1.GetListPostsRequest) (*post_v1.GetListPostsResponse, error) {
	if request.GetPage() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Page must be specified")
	}

	if request.GetPageSize() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Page Size must be specified")
	}

	res, err := sapi.repo.GetListPosts(ctx, int(request.GetPage()), int(request.GetPageSize()))
	if err != nil {
		return nil, err
	}

	// Convert res from []*models.Post to []*post_v1.PostResponse
	postResponses := make([]*post_v1.PostResponse, len(res))
	for i, post := range res {
		postResponses[i] = &post_v1.PostResponse{
			PostId:   post.ID,       // Ensure the ID type is compatible; convert if necessary
			AuthorId: post.AuthorID, // Ensure the AuthorID type is compatible; convert if necessary
			Title:    post.Title,
			Content:  post.Content,
		}
	}

	fmt.Print("The post list was goten successfully!")
	return &post_v1.GetListPostsResponse{
		Posts:    postResponses,
		NextPage: request.GetPage() + 1,
	}, nil
}
