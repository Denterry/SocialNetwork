package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Denterry/SocialNetwork/authService/pkg/post_v1"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
)

type PostServiceClientAPI struct {
	Client post_v1.PostServiceClient
}

func NewPostServiceClientAPI(client post_v1.PostServiceClient) *PostServiceClientAPI {
	return &PostServiceClientAPI{Client: client}
}

// TODO: Create a post
func (h *PostServiceClientAPI) CreatePost(c *gin.Context) {
	var req post_v1.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.Client.CreatePost(context.Background(), &req)
	if err != nil {
		grpcErr := status.Convert(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": grpcErr.Message()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// TODO: Update a post
func (h *PostServiceClientAPI) UpdatePost(c *gin.Context) {
	var req post_v1.UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// verify the user's permission to update the post here
	// req.AuthorId = extractUserId(c)

	res, err := h.Client.UpdatePost(context.Background(), &req)
	if err != nil {
		grpcErr := status.Convert(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": grpcErr.Message()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// TODO: Delete a post
func (h *PostServiceClientAPI) DeletePost(c *gin.Context) {
	// postId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
	// 	return
	// }

	var req post_v1.DeletePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// req := post_v1.DeletePostRequest{
	// 	PostId:   postId,
	// 	AuthorId: req.AuthorId, // verify the user is authorized to delete this post -> extractUserId(c)
	// }

	_, err := h.Client.DeletePost(context.Background(), &req)
	if err != nil {
		grpcErr := status.Convert(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": grpcErr.Message()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "post deleted"})
}

// TODO: Get a post
func (h *PostServiceClientAPI) GetPost(c *gin.Context) {
	postId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	res, err := h.Client.GetPostById(context.Background(), &post_v1.GetPostByIdRequest{
		PostId:   postId,
		AuthorId: 1, // verify the user is authorized to get this post
	})

	if err != nil {
		grpcErr := status.Convert(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": grpcErr.Message()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// TODO: Get a list of posts
func (h *PostServiceClientAPI) GetListPosts(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil || pageSize < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	req := post_v1.ListPostsRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		AuthorId: 1, // verify the user is authorized to get list of posts
	}

	res, err := h.Client.ListPosts(context.Background(), &req)
	if err != nil {
		grpcErr := status.Convert(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": grpcErr.Message()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// func extractUserId(c *gin.Context) int64 {
// 	// Extract user ID from JWT token or context
// 	_ = c
// 	return 12345
// }
