package server

import (
	"context"
	"fmt"
	"goplayground/pkg/grpc1/proto/sample"
	"net"

	"google.golang.org/grpc"
)

type sampleServer struct {
	sample.UnimplementedRunnerServer
}

func (s *sampleServer) Do(ctx context.Context, req *sample.DoRequest) (*sample.DoResponse, error) {
	out := sample.DoResponse{Item: req.Item}

	fmt.Println(req.Item)

	return &out, nil
}

func Server() error {
	s := grpc.NewServer()
	sample.RegisterRunnerServer(s, &sampleServer{})

	// Create a network listener
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		return err
	}

	// Serve GRPC client
	return s.Serve(lis)
}
