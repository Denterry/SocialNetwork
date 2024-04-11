package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/Denterry/SocialNetwork/postService/internal/config"
	"github.com/Denterry/SocialNetwork/postService/internal/grpc/server"
	"github.com/Denterry/SocialNetwork/postService/internal/storage"
	"github.com/Denterry/SocialNetwork/postService/internal/storage/postgres"
	"github.com/Denterry/SocialNetwork/postService/pkg/post_v1"
	"google.golang.org/grpc"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

// разобраться с подключаением к бд
// разобраться с обработкой чуваков тех, кто может крадить посты
// разобраться с докер компоузом
// разобраться с конфигом

func main() {
	cfg := config.MustLoad()

	postgres.InitDb()
	defer postgres.Db.Close()

	log := setupLogger(cfg.Env)
	log.Info("starting application", slog.Any("config", cfg))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		log.Error("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	repo := storage.NewPostRepositoryPg(postgres.Db)
	postService := server.NewServerAPI(repo)

	post_v1.RegisterPostServiceServer(grpcServer, postService)

	log.Info(fmt.Sprintf("Server is running on localhost:%d", cfg.GRPC.Port))
	if err := grpcServer.Serve(lis); err != nil {
		log.Error("Failed to serve: %v", err)
	}
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
