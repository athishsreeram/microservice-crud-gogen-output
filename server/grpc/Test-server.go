package grpc

import (
	"Test-output/proto"
	v1 "Test-output/service/v1"
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunServer(ctx context.Context, port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	v1API := v1.NewTestServiceServer()

	proto.RegisterTestServiceServer(srv, v1API)
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}
