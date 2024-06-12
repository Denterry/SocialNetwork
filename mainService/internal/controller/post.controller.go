package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Denterry/SocialNetwork/mainService/internal/config"
	"github.com/Denterry/SocialNetwork/mainService/internal/kafka"
	"github.com/Denterry/SocialNetwork/mainService/middleware"
	"github.com/Denterry/SocialNetwork/postService/pkg/post_v1"
	"github.com/Denterry/SocialNetwork/statisticsService/pkg/stat_v1"
	"github.com/gin-gonic/gin"
)

type PostController interface {
	CreatePost(ctx *gin.Context)
	UpdatePost(ctx *gin.Context)
	DeletePost(ctx *gin.Context)
	GetPost(ctx *gin.Context)
	GetListPosts(ctx *gin.Context)
	ViewPost(ctx *gin.Context)
	LikePost(ctx *gin.Context)
	GetStatPost(ctx *gin.Context)
	GetPostTop(ctx *gin.Context)
}

type postController struct {
	PostCLient       post_v1.PostServiceClient
	Cfg              *config.Config
	KafkaProducer    *kafka.KafkaProducer
	StatisticsClient stat_v1.StatisticsServiceClient
}

func NewPostController(engine *gin.Engine,
	postClient post_v1.PostServiceClient,
	cfg *config.Config,
	kafkaProducer *kafka.KafkaProducer,
	statisticsClient stat_v1.StatisticsServiceClient) {

	controller := &postController{
		PostCLient:       postClient,
		Cfg:              cfg,
		KafkaProducer:    kafkaProducer,
		StatisticsClient: statisticsClient,
	}

	api_protected := engine.Group("api/admin")
	{
		api_protected.Use(middleware.JWTAuthMiddleware(cfg))
		api_protected.POST("posts", controller.CreatePost)
		api_protected.PUT("posts/:id", controller.UpdatePost)
		api_protected.DELETE("posts/:id", controller.DeletePost)
		api_protected.GET("posts/:id", controller.GetPost)
		api_protected.GET("posts", controller.GetListPosts)
		api_protected.POST("posts/:id/view", controller.ViewPost)
		api_protected.POST("posts/:id/like", controller.LikePost)
		api_protected.GET("posts/:id/statistics", controller.GetStatPost)
		api_protected.GET("posts/topn", controller.GetPostTop)
	}
}

func (controller postController) GetPostTop(ctx *gin.Context) {
	request := &stat_v1.TopNPostsRequest{}

	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
		})
		return
	}

	res, err := controller.StatisticsClient.TopNPosts(ctx, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (controller postController) GetStatPost(ctx *gin.Context) {
	postID := ctx.Param("id")
	request := &stat_v1.TotalViewsLikesRequest{}

	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
		})
		return
	}

	if request.GetPostId() == 0 {
		pid, err := strconv.ParseInt(postID, 10, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		request.PostId = pid
	}

	res, err := controller.StatisticsClient.TotalViewsLikes(ctx, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (controller postController) LikePost(ctx *gin.Context) {
	postID := ctx.Param("id")
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	go func() {
		message := fmt.Sprintf(`{"postID": "%s", "userID": "%s", "event": "like"}`, postID, userID)
		if err := controller.KafkaProducer.SendMessage("post_events", message); err != nil {
			fmt.Printf("Failed to send like event: %v\n", err)
		}
	}()

	ctx.JSON(http.StatusOK, gin.H{"message": "Like event was sent"})
}

func (controller postController) ViewPost(ctx *gin.Context) {
	postID := ctx.Param("id")
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	go func() {
		message := fmt.Sprintf(`{"postID": "%s", "userID": "%s", "event": "view"}`, postID, userID)
		if err := controller.KafkaProducer.SendMessage("post_events", message); err != nil {
			fmt.Printf("Failed to send view event: %v\n", err)
		}
	}()

	ctx.JSON(http.StatusOK, gin.H{"message": "View event was sent"})
}

func (controller postController) CreatePost(ctx *gin.Context) {
	request := &post_v1.CreatePostRequest{}

	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
		})
		return
	}

	res, err := controller.PostCLient.CreatePost(context.Background(), request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (controller postController) UpdatePost(ctx *gin.Context) {
	request := &post_v1.UpdatePostRequest{}

	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
		})
		return
	}

	res, err := controller.PostCLient.UpdatePost(context.Background(), request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller postController) DeletePost(ctx *gin.Context) {
	request := &post_v1.DeletePostRequest{}

	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
		})
		return
	}

	res, err := controller.PostCLient.DeletePost(context.Background(), request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller postController) GetPost(ctx *gin.Context) {
	request := &post_v1.GetPostByIdRequest{}

	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
		})
		return
	}

	res, err := controller.PostCLient.GetPostById(context.Background(), request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller postController) GetListPosts(ctx *gin.Context) {
	request := &post_v1.GetListPostsRequest{}

	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
		})
		return
	}

	res, err := controller.PostCLient.GetListPosts(context.Background(), request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
