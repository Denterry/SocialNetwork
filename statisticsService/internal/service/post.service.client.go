package service

import (
	"log"

	"github.com/Denterry/SocialNetwork/postService/pkg/post_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewPostServiceClient(address string) post_v1.PostServiceClient {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to %s: %v", address, err)
	}

	return post_v1.NewPostServiceClient(conn)
}
