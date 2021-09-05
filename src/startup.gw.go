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
	endPoint := "localhost:8081"
	opt := []grpc.DialOption{grpc.WithInsecure(), grpc.WithReturnConnectionError()}

	err := pbs.RegisterCarwlerServiceHandlerFromEndpoint(context.Background(), gwmux, endPoint, opt)
	if err != nil {
		log.Fatal(err)
	}

	err = pbs.RegisterOCRServiceHandlerFromEndpoint(context.Background(), gwmux, endPoint, opt)
	if err != nil {
		log.Fatal(err)
	}

	httpServer := &http.Server{
		Addr:    ":9080",
		Handler: gwmux,
	}

	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
