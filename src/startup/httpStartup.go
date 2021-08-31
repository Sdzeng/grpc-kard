package main

import (
	"context"
	"grpc-kard/src/services/pbs"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {

	gwmux := runtime.NewServeMux()
	opt := []grpc.DialOption{}
	err := pbs.RegisterCarwlerServiceHandlerFromEndpoint(context.Background(), gwmux, "localhost:8081", opt)
	if err != nil {
		log.Fatal(err)
	}

	httpServer := &http.Server{Addr: "8080", Handler: gwmux}

	httpServer.ListenAndServe()
}
