package service

import (
	"context"
	"log"

	"github.com/Denterry/SocialNetwork/postService/pkg/post_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:generate mockgen -source=post.service.client.go -destination=mocks/post.mock.go

type PostServiceClient interface {
	// Creates a new post
	CreatePost(ctx context.Context, in *post_v1.CreatePostRequest, opts ...grpc.CallOption) (*post_v1.PostResponse, error)
	// Updates an existing post
	UpdatePost(ctx context.Context, in *post_v1.UpdatePostRequest, opts ...grpc.CallOption) (*post_v1.PostResponse, error)
	// Deletes a specific post
	DeletePost(ctx context.Context, in *post_v1.DeletePostRequest, opts ...grpc.CallOption) (*post_v1.PostResponse, error)
	// Retrieves a specific post by ID
	GetPostById(ctx context.Context, in *post_v1.GetPostByIdRequest, opts ...grpc.CallOption) (*post_v1.PostResponse, error)
	// Lists posts with pagination
	GetListPosts(ctx context.Context, in *post_v1.GetListPostsRequest, opts ...grpc.CallOption) (*post_v1.GetListPostsResponse, error)
}

func NewPostServiceClient(address string) PostServiceClient {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to %s: %v", address, err)
	}

	return post_v1.NewPostServiceClient(conn)
}
