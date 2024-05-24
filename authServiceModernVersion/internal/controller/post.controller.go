package controller

import (
	"context"
	"net/http"

	"github.com/Denterry/SocialNetwork/postService/pkg/post_v1"
	"github.com/gin-gonic/gin"
)

type PostController interface {
	CreatePost(ctx *gin.Context)
	UpdatePost(ctx *gin.Context)
	DeletePost(ctx *gin.Context)
	GetPost(ctx *gin.Context)
	GetListPosts(ctx *gin.Context)
}

type postController struct {
	CLient post_v1.PostServiceClient
}

func NewPostController(engine *gin.Engine, client post_v1.PostServiceClient) {
	controller := &postController{
		CLient: client,
	}

	api := engine.Group("api")
	{
		api.POST("posts", controller.CreatePost)
		api.PUT("posts/:id", controller.UpdatePost)
		api.DELETE("posts/:id", controller.DeletePost)
		api.GET("posts/:id", controller.GetPost)
		api.GET("posts", controller.GetListPosts)
	}
}

func (controller postController) CreatePost(ctx *gin.Context) {
	request := &post_v1.CreatePostRequest{}

	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
		})
		return
	}

	res, err := controller.CLient.CreatePost(context.Background(), request)
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

	res, err := controller.CLient.UpdatePost(context.Background(), request)
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

	res, err := controller.CLient.DeletePost(context.Background(), request)
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

	res, err := controller.CLient.GetPostById(context.Background(), request)
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

	res, err := controller.CLient.GetListPosts(context.Background(), request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
