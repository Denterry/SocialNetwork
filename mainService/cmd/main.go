package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Denterry/SocialNetwork/mainService/domain"
	"github.com/Denterry/SocialNetwork/mainService/internal/config"
	"github.com/Denterry/SocialNetwork/mainService/internal/controller"
	"github.com/Denterry/SocialNetwork/mainService/internal/kafka"
	"github.com/Denterry/SocialNetwork/mainService/internal/repository"
	"github.com/Denterry/SocialNetwork/mainService/internal/service"
	"github.com/Denterry/SocialNetwork/mainService/internal/storage"
	"github.com/gin-gonic/gin"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	db, err := storage.InitialiseDB(&storage.DbConfig{
		User:     cfg.Storage.User,
		Password: cfg.Storage.Password,
		DbName:   cfg.Storage.Name,
		Host:     cfg.Storage.Host,
		Port:     cfg.Storage.Port,
		Schema:   cfg.Storage.Schema,
	})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		panic(err)
	}

	log := setupLogger(cfg.Env)
	log.Info("starting Authentication application", slog.Any("config", cfg))

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, cfg)

	postServiceClient := service.NewPostServiceClient(fmt.Sprintf("%s:%s", cfg.PostService.Host, cfg.PostService.Port))

	// Kafka Action
	kafkProducer, err := kafka.NewKafkaProducer(cfg)
	if err != nil {
		log.Error("Error while running NewKafkaProducer:", err)
	}

	// API Action
	engine := gin.Default()

	controller.NewUserController(engine, userService, cfg)
	controller.NewPostController(engine, postServiceClient, cfg, kafkProducer)

	err = engine.Run(fmt.Sprintf("%s:%s", cfg.Gin.Host, cfg.Gin.Port))
	if err != nil {
		log.Info(fmt.Sprintf("Server running is failed on %s:%s", cfg.Gin.Host, cfg.Gin.Port))
	}

	log.Info(fmt.Sprintf("Server is running on %s:%s", cfg.Gin.Host, cfg.Gin.Port))
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
