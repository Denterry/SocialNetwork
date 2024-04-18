package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/Denterry/SocialNetwork/postService/app/grpcApp"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func NewApp(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// TODO: firstly we need ti initialize our storage

	// TODO:  init gprc server (server ps)

	grpcApp := grpcapp.NewApp(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
