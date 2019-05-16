package rest

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	gw "Test-output/proto"
)

func RunServer(ctx context.Context, grpcport string, httpport string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterTestServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcport, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":"+httpport, mux)
}
