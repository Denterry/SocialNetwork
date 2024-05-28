package controller

import (
	"github.com/Denterry/SocialNetwork/statisticsService/internal/config"
	"github.com/gin-gonic/gin"
)

type EventController interface {
	HealthCheck(ctx *gin.Context)
}

type eventController struct {
	cfg *config.Config
}

func NewEventController(engine *gin.Engine, cfg *config.Config) {
	controller := &eventController{
		cfg: cfg,
	}

	api := engine.Group("api")
	{
		api.POST("check", controller.HealthCheck)
	}
}

func (controller *eventController) HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "WE ARE ALIVE!!!!!"})
}
