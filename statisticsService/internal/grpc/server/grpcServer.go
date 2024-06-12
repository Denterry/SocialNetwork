package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/Denterry/SocialNetwork/postService/pkg/post_v1"
	"github.com/Denterry/SocialNetwork/statisticsService/internal/config"
	"github.com/Denterry/SocialNetwork/statisticsService/internal/domain/models"
	"github.com/Denterry/SocialNetwork/statisticsService/pkg/stat_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// StatRepository provides an interface for CRUD operations on statistics information in Clickhouse
type StatRepository interface {
	GetLikesViewsOnPost(context.Context, int64) (uint64, uint64, error)
	GetTopNPosts(context.Context, int64, string) ([]*models.PostTop, error)
	GetAllPostsWithLikes(ctx context.Context) ([]*models.PostTop, error)
}

type serverAPI struct {
	stat_v1.StatisticsServiceServer
	repo              StatRepository
	postServiceClient post_v1.PostServiceClient
	cfg               *config.Config
}

func NewServerAPI(repo StatRepository, psc post_v1.PostServiceClient, cfg *config.Config) *serverAPI {
	return &serverAPI{
		repo:              repo,
		postServiceClient: psc,
		cfg:               cfg,
	}
}

func (sapi *serverAPI) TotalViewsLikes(ctx context.Context, request *stat_v1.TotalViewsLikesRequest) (*stat_v1.TotalViewsLikesResponse, error) {
	if request.GetPostId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "Post Id must be specified")
	}

	totalViews, totalLikes, err := sapi.repo.GetLikesViewsOnPost(ctx, request.GetPostId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &stat_v1.TotalViewsLikesResponse{
		PostId:     request.GetPostId(),
		TotalViews: totalViews,
		TotalLikes: totalLikes,
	}, nil
}

func (sapi *serverAPI) TopNPosts(ctx context.Context, request *stat_v1.TopNPostsRequest) (*stat_v1.TopNPostsResponse, error) {
	var n int64
	if request.GetN() == 0 {
		n = 5
	} else {
		n = request.GetN()
	}

	if request.GetEvent() == "" {
		return nil, status.Error(codes.InvalidArgument, "Event type must be specified")
	}

	topNPosts, err := sapi.repo.GetTopNPosts(ctx, n, request.GetEvent())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var top []*stat_v1.TopNPosts
	for _, post := range topNPosts {

		// Have to get authorId from post information that contains in post service
		postInfo, err := sapi.postServiceClient.GetPostById(ctx, &post_v1.GetPostByIdRequest{
			PostId:   int64(post.PostId),
			AuthorId: "",
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		// Have to get real username from main service
		url := fmt.Sprintf("http://%s:%s/api/users/%s", sapi.cfg.MainService.Host, sapi.cfg.MainService.Port, postInfo.AuthorId)

		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return nil, err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to fetch author details: %s", resp.Status)
		}

		var msResponse models.MSResponse
		if err := json.NewDecoder(resp.Body).Decode(&msResponse); err != nil {
			return nil, err
		}

		top = append(top, &stat_v1.TopNPosts{
			PostId:   post.PostId,
			UserId:   postInfo.AuthorId,
			Username: msResponse.Data.Username,
			Number:   post.Number,
		})
	}

	return &stat_v1.TopNPostsResponse{
		TopNPosts: top,
	}, nil
}

func (sapi *serverAPI) TopNUsers(ctx context.Context, request *stat_v1.TopNUsersRequest) (*stat_v1.TopNUsersResponse, error) {
	var n int64
	if request.GetN() == 0 {
		n = 3
	} else {
		n = request.GetN()
	}

	postsWithLikes, err := sapi.repo.GetAllPostsWithLikes(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	userTop := make(map[string]uint64)
	for _, post := range postsWithLikes {
		// Have to get authorId from post information that contains in post service
		postInfo, err := sapi.postServiceClient.GetPostById(ctx, &post_v1.GetPostByIdRequest{
			PostId:   int64(post.PostId),
			AuthorId: "",
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		// Have to get real username from main service
		url := fmt.Sprintf("http://%s:%s/api/users/%s", sapi.cfg.MainService.Host, sapi.cfg.MainService.Port, postInfo.AuthorId)

		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return nil, err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to fetch author details: %s", resp.Status)
		}

		var msResponse models.MSResponse
		if err := json.NewDecoder(resp.Body).Decode(&msResponse); err != nil {
			return nil, err
		}

		userTop[msResponse.Data.Username] += post.Number
	}

	top := make([]*stat_v1.TopNUsers, 0, n)
	for k, v := range userTop {
		top = append(top, &stat_v1.TopNUsers{
			Username: k,
			Number:   v,
		})
	}

	sort.Slice(top, func(i, j int) bool {
		return top[i].Number > top[j].Number
	})

	if int64(len(top)) < n {
		n = int64(len(top))
	}
	return &stat_v1.TopNUsersResponse{
		TopNUsers: top[:n],
	}, nil
}
