package main

import (
	"log"
	"net"

	"github.com/Denterry/SocialNetwork/postService/api/proto"
	"github.com/Denterry/SocialNetwork/postService/internal/storage"
	postdealer "github.com/Denterry/SocialNetwork/postService/pkg/postDealer"
	"google.golang.org/grpc"
)

func main() {
	storage.InitDb()
	defer storage.Db.Close()

	lis, err := net.Listen("tcp", ":5005")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	srv := &postdealer.GRPCServer{}
	proto.RegisterPostServiceServer(grpcServer, srv)

	postRepository := repository.NewPostRepository()
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	pb.RegisterPostServiceServer(grpcServer, postHandler)

	log.Println("gRPC server is running...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
