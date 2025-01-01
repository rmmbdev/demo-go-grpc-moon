package server

import (
	"github.com/rmmbdev/demo-go-grpc-moon.git/gen/apps/gateway/pb"
	"log"

	"golang.org/x/net/context"
)

type Server struct {
	pb.UnimplementedChatServiceServer
}

func (s *Server) GetMessage(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &pb.Response{Body: "Hello From the Server!"}, nil
}
