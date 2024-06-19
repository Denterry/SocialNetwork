package server

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Denterry/SocialNetwork/postService/pkg/post_v1"
	"github.com/Denterry/SocialNetwork/statisticsService/internal/config"
	"github.com/Denterry/SocialNetwork/statisticsService/internal/domain/models"
	mock_server "github.com/Denterry/SocialNetwork/statisticsService/internal/grpc/server/mocks"
	mock_service "github.com/Denterry/SocialNetwork/statisticsService/internal/service/mocks"
	"github.com/Denterry/SocialNetwork/statisticsService/pkg/stat_v1"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestHandler_totalViewsLikes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_server.NewMockStatRepository(ctrl)
	s := &serverAPI{repo: mockRepo}

	postID := int64(123)
	likes := uint64(10)
	views := uint64(100)

	mockRepo.EXPECT().GetLikesViewsOnPost(gomock.Any(), postID).Return(views, likes, nil)

	req := &stat_v1.TotalViewsLikesRequest{PostId: postID}
	resp, err := s.TotalViewsLikes(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, postID, resp.PostId)
	require.Equal(t, views, resp.TotalViews)
	require.Equal(t, likes, resp.TotalLikes)
}

func TestHandler_topNPosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_server.NewMockStatRepository(ctrl)
	mockPostServiceClient := mock_service.NewMockPostServiceClient(ctrl)
	s := &serverAPI{repo: mockRepo, postServiceClient: mockPostServiceClient}

	event := "like"
	n := int64(2)
	post1 := &models.PostTop{PostId: 1,
		Number: 10,
	}
	post2 := &models.PostTop{PostId: 2,
		Number: 20,
	}

	mockRepo.EXPECT().GetTopNPosts(gomock.Any(), n, event).Return([]*models.PostTop{post1, post2}, nil)

	mockPostServiceClient.EXPECT().GetPostById(gomock.Any(), &post_v1.GetPostByIdRequest{PostId: 1, AuthorId: ""}).Return(&post_v1.PostResponse{PostId: 1, Title: "Test Title 1", Content: "Test Content  1", AuthorId: "uuid-uuid-uuid-uuid-uuid1"}, nil)
	mockPostServiceClient.EXPECT().GetPostById(gomock.Any(), &post_v1.GetPostByIdRequest{PostId: 2, AuthorId: ""}).Return(&post_v1.PostResponse{PostId: 2, Title: "Test Title 2", Content: "Test Content  2", AuthorId: "uuid-uuid-uuid-uuid-uuid2"}, nil)

	// Mocking HTTP client
	authorResponse1 := models.MSResponse{Data: models.AuthorInfo{Username: "user1"}}
	authorResponse2 := models.MSResponse{Data: models.AuthorInfo{Username: "user2"}}

	author1Server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(authorResponse1)
	}))
	defer author1Server.Close()

	author2Server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(authorResponse2)
	}))
	defer author2Server.Close()

	// Parse the URL to get the host and port
	parsedURL, err := url.Parse(author1Server.URL)
	require.NoError(t, err)
	host := parsedURL.Hostname()
	port := parsedURL.Port()

	s.cfg = &config.Config{
		MainService: config.MainServiceConfig{
			Host: host,
			Port: port,
			URL:  author1Server.URL,
		},
	}

	req := &stat_v1.TopNPostsRequest{N: n, Event: event}
	resp, err := s.TopNPosts(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Len(t, resp.TopNPosts, 2)

	// Validate the first post
	require.Equal(t, uint32(0x1), resp.TopNPosts[0].PostId)
	require.Equal(t, "uuid-uuid-uuid-uuid-uuid1", resp.TopNPosts[0].UserId)
	require.Equal(t, "user1", resp.TopNPosts[0].Username)
	require.Equal(t, uint64(10), resp.TopNPosts[0].Number)

	// Validate the second post
	require.Equal(t, uint32(0x2), resp.TopNPosts[1].PostId)
	require.Equal(t, "uuid-uuid-uuid-uuid-uuid2", resp.TopNPosts[1].UserId)
	require.Equal(t, "user1", resp.TopNPosts[1].Username)
	require.Equal(t, uint64(20), resp.TopNPosts[1].Number)
}
