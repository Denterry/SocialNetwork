package service

import (
	"log"

	"github.com/Denterry/SocialNetwork/statisticsService/pkg/stat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewStatServiceClient(address string) stat_v1.StatisticsServiceClient {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to %s: %v", address, err)
	}

	return stat_v1.NewStatisticsServiceClient(conn)
}
