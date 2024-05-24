package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Denterry/SocialNetwork/authServiceModernVersion/domain"
	"github.com/Denterry/SocialNetwork/authServiceModernVersion/internal/config"
	"github.com/Denterry/SocialNetwork/authServiceModernVersion/internal/controller"
	"github.com/Denterry/SocialNetwork/authServiceModernVersion/internal/repository"
	"github.com/Denterry/SocialNetwork/authServiceModernVersion/internal/service"
	"github.com/Denterry/SocialNetwork/authServiceModernVersion/internal/storage"
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
	userService := service.NewUserService(userRepository)

	postServiceClient := service.NewPostServiceClient(fmt.Sprintf("%s:%s", cfg.PostService.Host, cfg.PostService.Port))

	engine := gin.Default()

	controller.NewUserController(engine, userService)
	controller.NewPostController(engine, postServiceClient)

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
