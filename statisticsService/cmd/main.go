package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Denterry/SocialNetwork/statisticsService/internal/config"
	"github.com/Denterry/SocialNetwork/statisticsService/internal/controller"
	"github.com/Denterry/SocialNetwork/statisticsService/internal/kafka"
	"github.com/Denterry/SocialNetwork/statisticsService/internal/storage"
	"github.com/gin-gonic/gin"
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
	_ = conn

	fmt.Println("asdaskdlkaghsdhjglakfgahs hjasvf ashyf hjasvlhj")

	// Kafka Action
	kafkaConsumer, err := kafka.NewKafkaConsumer(cfg)
	if err != nil {
		log.Error("Failed to create kafka consumer: ", err)
	}
	log.Info("Successfully create new kafka consumer")

	go kafkaConsumer.ConsumeEvents(conn)

	// API Action
	engine := gin.Default()

	controller.NewEventController(engine, cfg)

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
