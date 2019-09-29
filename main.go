package main

import (
	"context" // Use "golang.org/x/net/context" for Golang version <= 1.6
	"flag"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	config "github.com/u03013112/ss-pb/config"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "config:50001", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := config.RegisterSSConfigHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Println("listen 8081")
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
