package client

import (
	"log"
	"time"

	"github.com/Denterry/SocialNetwork/authService/pkg/post_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewPostServiceClient(address string) post_v1.PostServiceClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Fatalf("could not connect to %s: %v", address, err)
	}
	return post_v1.NewPostServiceClient(conn)
}
