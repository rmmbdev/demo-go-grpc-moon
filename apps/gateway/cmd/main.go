package main

import (
	"github.com/rmmbdev/demo-go-grpc-moon.git/apps/gateway/internal/server"
	"github.com/rmmbdev/demo-go-grpc-moon.git/gen/apps/gateway/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	log.Printf("gateway server listening at %v", lis.Addr())

	chat := server.Server{}
	pb.RegisterChatServiceServer(grpcServer, &chat)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
