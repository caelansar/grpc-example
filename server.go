package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"rpcTestProj/hello"
)

const (
	port = ":50051"
)

type server struct{}

// implement interface
func (s *server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	hello.RegisterGreeterServer(s, &server{})
	reflection.Register(s) // register server in rpc
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}