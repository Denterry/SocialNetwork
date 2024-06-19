package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"sync"

	"github.com/Denterry/SocialNetwork/statisticsService/internal/config"
	"github.com/Denterry/SocialNetwork/statisticsService/internal/controller"
	"github.com/Denterry/SocialNetwork/statisticsService/internal/grpc/server"
	"github.com/Denterry/SocialNetwork/statisticsService/internal/kafka"
	"github.com/Denterry/SocialNetwork/statisticsService/internal/repository"
	"github.com/Denterry/SocialNetwork/statisticsService/internal/service"
	"github.com/Denterry/SocialNetwork/statisticsService/internal/storage"
	"github.com/Denterry/SocialNetwork/statisticsService/pkg/stat_v1"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("starting Authentication application", slog.Any("config", cfg))

	conn, err := storage.InitialiseDB(&storage.DbConfig{
		User:     cfg.Storage.User,
		Password: "",
		DbName:   cfg.Storage.Name,
		Host:     cfg.Storage.Host,
		Port:     cfg.Storage.Port,
	})
	if err != nil {
		log.Error("Failed to connect to ClickHouse: ", err)
	}

	// 	migrationUp := `
	// 	-- Insert test events
	// 	INSERT INTO post_events (postID, userID, event) VALUES
	// 		(1, '6e3fd317-05f2-4e5c-bcf0-b270119e3fea', 'like'),
	// 		(1, '6e3fd317-05f2-4e5c-bcf0-b270119e3fea', 'view');
	// `
	// 	// Execute migration
	// 	err = conn.Exec(context.Background(), migrationUp)
	// 	if err != nil {
	// 		log.Error("Failed to execute migration: ", err)
	// 		return
	// 	}

	// 	log.Info("Migration applied successfully!")

	var wg sync.WaitGroup
	wg.Add(3)

	// Kafka Action
	kafkaConsumer, err := kafka.NewKafkaConsumer(cfg)
	if err != nil {
		log.Error("Failed to create kafka consumer: ", err)
	}
	log.Info("Successfully create new kafka consumer")

	go func() {
		defer wg.Done()
		kafkaConsumer.ConsumeEvents(conn)
	}()

	// gRPC Action
	go func() {
		defer wg.Done()
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.GRPC.Host, cfg.GRPC.Port))
		if err != nil {
			log.Error("Failed to listen: %v", err)
		}

		postServiceClient := service.NewPostServiceClient(fmt.Sprintf("%s:%s", cfg.PostService.Host, cfg.PostService.Port))

		grpcServer := grpc.NewServer()
		repo := repository.NewStatRepositoryClickhouse(conn)
		statService := server.NewServerAPI(repo, postServiceClient, cfg)

		stat_v1.RegisterStatisticsServiceServer(grpcServer, statService)

		log.Info(fmt.Sprintf("gRPC Server is running on %s:%s", cfg.GRPC.Host, cfg.GRPC.Port))
		if err := grpcServer.Serve(lis); err != nil {
			log.Error("Failed to serve: %v", err)
		}
	}()

	// API Action
	go func() {
		defer wg.Done()
		engine := gin.Default()

		controller.NewEventController(engine, cfg)

		err = engine.Run(fmt.Sprintf("%s:%s", cfg.Gin.Host, cfg.Gin.Port))

		if err != nil {
			log.Info(fmt.Sprintf("API Server running is failed on %s:%s", cfg.Gin.Host, cfg.Gin.Port))
		}

		log.Info(fmt.Sprintf("API Server is running on %s:%s", cfg.Gin.Host, cfg.Gin.Port))
	}()

	wg.Wait()
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
