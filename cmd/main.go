package main

import (
	"net"

	"github.com/ArtuoS/hurb/cmd/server"
	"github.com/ArtuoS/hurb/pb"
	"google.golang.org/grpc"
)

func main() {
	println("Running gRPC")

	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterCurrencyServer(s, &server.Server{})
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
