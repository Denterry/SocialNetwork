package controller

import (
	"net/http"

	"github.com/Denterry/SocialNetwork/authServiceModernVersion/internal/service"
	"github.com/Denterry/SocialNetwork/authServiceModernVersion/model"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Signup(ctx *gin.Context)
	ChangeInfo(ctx *gin.Context)
	Signin(ctx *gin.Context)
}

type userController struct {
	service service.UserService
}

func NewUserController(engine *gin.Engine, userService service.UserService) {
	controller := &userController{
		service: userService,
	}

	api := engine.Group("api")
	{
		api.POST("users/sign-up", controller.Signup)
		api.PUT("users/change-info", controller.ChangeInfo)
		api.POST("users/sign-in", controller.Signin)
	}
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

	err := controller.service.Signin(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User authenticated successfully"})
}
