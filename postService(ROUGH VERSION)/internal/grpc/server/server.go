package server

import (
	"context"
	"strconv"

	psv1 "github.com/Denterry/SocialNetwork/postService/api/protos/gen/go/ps"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PostResp struct {
	PostId   int64
	Title    string
	Content  string
	AuthorId string
}

type PostsResp struct {
	Posts []*PostResp
}

type PostService interface {
	CreatePost(
		ctx context.Context,
		authorId int64,
		title string,
		content string,
	) (PostId int64, Title string, Content string, AuthorId string, err error)
	UpdatePost(
		ctx context.Context,
		postrId int64,
		title string,
		content string,
	) (PostId int64, Title string, Content string, AuthorId string, err error)
	DeletePost(
		ctx context.Context,
		postrId int64,
	) (PostId int64, Title string, Content string, AuthorId string, err error)
	GetPostById(
		ctx context.Context,
		postrId int64,
	) (PostId int64, Title string, Content string, AuthorId string, err error)
	GetPosts(
		ctx context.Context,
		page int32,
		pageSize int32,
	) (Posts []*PostResp, err error)
}

type serverAPI struct {
	psv1.UnimplementedPostServiceServer
	postServ PostService
}

func Register(gRPC *grpc.Server, postServ PostService) {
	psv1.RegisterPostServiceServer(gRPC, &serverAPI{postServ: postServ})
}

// TODO: Implement methods for the API
func (a *serverAPI) CreatePost(ctx context.Context, request *psv1.CreatePostRequest) (*psv1.PostResponse, error) {
	if request.GetAuthorId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Author Id must be specified")
	}

	if request.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "Title must be specified")
	}

	if request.GetContent() == "" {
		return nil, status.Error(codes.InvalidArgument, "Content must be specified")
	}

	PostId, Title, Content, AuthorId, err := a.postServ.CreatePost(ctx, request.GetAuthorId(), request.GetTitle(), request.GetContent())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &psv1.PostResponse{
		PostId:   strconv.FormatInt(int64(PostId), 10),
		Title:    Title,
		Content:  Content,
		AuthorId: AuthorId,
	}, nil
}

func (a *serverAPI) UpdatePost(ctx context.Context, request *psv1.UpdatePostRequest) (*psv1.PostResponse, error) {
	if request.GetPostId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Author Id must be specified")
	}

	if request.GetTitle() == "" {
		return nil, status.Error(codes.InvalidArgument, "Title must be specified")
	}

	if request.GetContent() == "" {
		return nil, status.Error(codes.InvalidArgument, "Content must be specified")
	}

	PostId, Title, Content, AuthorId, err := a.postServ.UpdatePost(ctx, request.GetPostId(), request.GetTitle(), request.GetContent())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &psv1.PostResponse{
		PostId:   strconv.FormatInt(int64(PostId), 10),
		Title:    Title,
		Content:  Content,
		AuthorId: AuthorId,
	}, nil
}

func (a *serverAPI) DeletePost(ctx context.Context, request *psv1.DeletePostRequest) (*psv1.PostResponse, error) {
	if request.GetPostId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Author Id must be specified")
	}

	PostId, Title, Content, AuthorId, err := a.postServ.DeletePost(ctx, request.GetPostId())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &psv1.PostResponse{
		PostId:   strconv.FormatInt(int64(PostId), 10),
		Title:    Title,
		Content:  Content,
		AuthorId: AuthorId,
	}, nil
}

func (a *serverAPI) GetPostById(ctx context.Context, request *psv1.GetPostByIdRequest) (*psv1.PostResponse, error) {
	if request.GetPostId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Author Id must be specified")
	}

	PostId, Title, Content, AuthorId, err := a.postServ.GetPostById(ctx, request.GetPostId())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &psv1.PostResponse{
		PostId:   strconv.FormatInt(int64(PostId), 10),
		Title:    Title,
		Content:  Content,
		AuthorId: AuthorId,
	}, nil
}

func (a *serverAPI) GetPosts(ctx context.Context, request *psv1.GetPostsRequest) (*psv1.GetPostsResponse, error) {
	if request.GetPage() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Page must be specified")
	}

	if request.GetPageSize() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Page Size must be specified")
	}

	Posts, err := a.postServ.GetPosts(ctx, request.GetPage(), request.GetPageSize())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &psv1.GetPostsResponse{
		Posts: Posts,
	}, nil
}
