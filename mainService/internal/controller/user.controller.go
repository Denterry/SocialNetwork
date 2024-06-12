package controller

import (
	"fmt"
	"net/http"

	"github.com/Denterry/SocialNetwork/mainService/internal/config"
	"github.com/Denterry/SocialNetwork/mainService/internal/service"
	"github.com/Denterry/SocialNetwork/mainService/middleware"
	"github.com/Denterry/SocialNetwork/mainService/model"
	"github.com/Denterry/SocialNetwork/mainService/util"
	"github.com/Denterry/SocialNetwork/statisticsService/pkg/stat_v1"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController interface {
	Signup(ctx *gin.Context)
	ChangeInfo(ctx *gin.Context)
	Signin(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
	CurrentUser(ctx *gin.Context)
	GetUserInfo(ctx *gin.Context)
	GetUserTop(ctx *gin.Context)
}

type userController struct {
	service           service.UserService
	cfg               *config.Config
	statisticsService stat_v1.StatisticsServiceClient
}

func NewUserController(engine *gin.Engine, userService service.UserService, cfg *config.Config, statisticsService stat_v1.StatisticsServiceClient) {
	controller := &userController{
		service:           userService,
		cfg:               cfg,
		statisticsService: statisticsService,
	}

	api := engine.Group("api/users")
	{
		api.POST("/sign-up", controller.Signup)
		api.POST("/sign-in", controller.Signin)
		api.POST("/retrieve", controller.Retrieve)
		api.GET("/:id", controller.GetUserInfo)
	}

	api_protected := engine.Group("api/admin")
	{
		api_protected.Use(middleware.JWTAuthMiddleware(cfg))
		api_protected.GET("user", controller.CurrentUser)
		api_protected.PUT("change-info", controller.ChangeInfo)
		api_protected.GET("top", controller.GetUserTop)
	}
}

func (controller userController) GetUserTop(ctx *gin.Context) {
	request := &stat_v1.TopNUsersRequest{}

	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
		})
		return
	}

	res, err := controller.statisticsService.TopNUsers(ctx, request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (controller userController) Signup(ctx *gin.Context) {
	request := &model.SignupRequest{}

	// The ShouldBind method inspects the incoming Content-Type header,
	// selects a strategy to deserialise the request body and then bind and
	// validate the parsed parameter map according to the binding tags if any,
	// present against the struct fields pointed by the pointer passed to it.
	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
		})
		return
	}

	err := controller.service.Signup(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "You have successfully registered!"})
}

func (controller userController) ChangeInfo(ctx *gin.Context) {
	request := &model.ChangeInfoRequest{}

	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
		})
		return
	}

	err := controller.service.ChangeInfo(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "You have successfully changed your personal info!"})
}

func (controller userController) Signin(ctx *gin.Context) {
	request := &model.SigninRequest{}

	if err := ctx.ShouldBindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
		})
		return
	}

	token, err := controller.service.Signin(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (controller *userController) Retrieve(ctx *gin.Context) {
	request := &model.RetrieveRequest{}
	if err := ctx.ShouldBind(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
		})
		return
	}

	user, err := controller.service.Retrieve(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (controller *userController) CurrentUser(ctx *gin.Context) {
	user_id, err := util.ExtractTokenID(ctx, controller.cfg)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := controller.service.CurrentUser(&model.UserIdRequest{
		UserID: user_id,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

func (controller *userController) GetUserInfo(ctx *gin.Context) {
	requestedUserId := ctx.Param("id")

	uid, err := uuid.Parse(requestedUserId) // Parse string to UUID
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("user id is not a valid UUID: %v", err),
		})
		return
	}

	userInfo, err := controller.service.GetUserInfo(&model.UserIdRequest{
		UserID: uid,
	})

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": userInfo})
}
