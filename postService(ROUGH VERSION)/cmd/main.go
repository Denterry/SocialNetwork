package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Denterry/SocialNetwork/postService/app"
	"github.com/Denterry/SocialNetwork/postService/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	// fmt.Println(cfg)

	log := setupLogger(cfg.Env)
	log.Info("starting application", slog.Any("config", cfg))
	// log.Info("starting application",
	// 	slog.String("env", cfg.Env),
	// 	slog.Any("cfg", cfg),
	// 	slog.Int("port", cfg.GRPC.Port),
	// )
	// log.Debug("debug message")
	// log.Error("error message")
	// log.Warn("warn message")

	application := app.NewApp(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	go application.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info("stopping application", slog.String("signal", sign.String()))

	application.GRPCSrv.Stop()

	log.Info("application stopped")
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

// func main() {
// 	storage.InitDb()
// 	defer storage.Db.Close()

// 	lis, err := net.Listen("tcp", ":5005")
// 	if err != nil {
// 		log.Fatalf("Failed to listen: %v", err)
// 	}

// 	grpcServer := grpc.NewServer()
// 	srv := &postdealer.GRPCServer{}
// 	proto.RegisterPostServiceServer(grpcServer, srv)

// 	postRepository := repository.NewPostRepository()
// 	postService := service.NewPostService(postRepository)
// 	postHandler := handler.NewPostHandler(postService)

// 	pb.RegisterPostServiceServer(grpcServer, postHandler)

// 	log.Println("gRPC server is running...")
// 	if err := grpcServer.Serve(lis); err != nil {
// 		log.Fatalf("Failed to serve: %v", err)
// 	}
// }
