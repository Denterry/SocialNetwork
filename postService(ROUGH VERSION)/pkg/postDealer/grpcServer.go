package postdealer

import (
	"context"

	pb "github.com/Denterry/SocialNetwork/postService/api/proto"
	"github.com/Denterry/SocialNetwork/postService/internal/storage"
	"github.com/Denterry/SocialNetwork/postService/pkg/models"
)

type GRPCServer struct {
	postRepository storage.PostRepository
}

// Realization of CreatePost method for PostServiceServer Interface
func (s *GRPCServer) CreatePost(ctx context.Context, request *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	post := &models.Post{
		UserID:  request.UserId,
		Title:   request.Title,
		Content: request.Content,
	}

	postID, err := s.postRepository.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}
	return &pb.CreatePostResponse{PostId: postID}, nil
}

// func (s *GRPCServer) UpdatePost(ctx context.Context, request *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
// 	// Реализация метода UpdatePost
// 	return nil, nil
// }

// func (s *GRPCServer) DeletePost(ctx context.Context, request *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
// 	// Реализация метода DeletePost
// 	return nil, nil
// }

// func (s *GRPCServer) GetPostById(ctx context.Context, request *pb.GetPostByIdRequest) (*pb.GetPostByIdResponse, error) {
// 	// Реализация метода GetPostById
// 	return nil, nil
// }

// func (s *GRPCServer) GetPosts(ctx context.Context, request *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
// 	// Реализация метода GetPosts
// 	return nil, nil
// }

// func NewPostService(postRepository repository.PostRepository) *GRPCServer {
// 	return &PostService{postRepository: postRepository}
// }
