package main

import (
	"grpc-kard/src/services"
	"grpc-kard/src/services/pbs"
	"net"

	"google.golang.org/grpc"
)

func main() {

	rpcService := grpc.NewServer()
	pbs.RegisterCarwlerServiceServer(rpcService, new(services.CrawlerService))
	lis, _ := net.Listen("tcp", ":8081")

	rpcService.Serve(lis)
}
