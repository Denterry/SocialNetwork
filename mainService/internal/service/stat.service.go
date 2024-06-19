package service

import (
	"context"
	"log"

	"github.com/Denterry/SocialNetwork/statisticsService/pkg/stat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:generate mockgen -source=stat.service.go -destination=mocks/stat.mock.go

type StatisticsServiceClient interface {
	// Returns the total number of views and likes for a specific post
	TotalViewsLikes(ctx context.Context, in *stat_v1.TotalViewsLikesRequest, opts ...grpc.CallOption) (*stat_v1.TotalViewsLikesResponse, error)
	// Returns the top N posts by number of likes or views
	TopNPosts(ctx context.Context, in *stat_v1.TopNPostsRequest, opts ...grpc.CallOption) (*stat_v1.TopNPostsResponse, error)
	// Returns the top N users with the most likes across all their posts.
	TopNUsers(ctx context.Context, in *stat_v1.TopNUsersRequest, opts ...grpc.CallOption) (*stat_v1.TopNUsersResponse, error)
}

func NewStatServiceClient(address string) StatisticsServiceClient {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to %s: %v", address, err)
	}

	return stat_v1.NewStatisticsServiceClient(conn)
}
