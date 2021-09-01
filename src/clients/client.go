package main

import (
	"context"
	"fmt"
	"grpc-kard/src/services/pbs"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8081")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	sarwlerClient := pbs.NewCarwlerServiceClient(conn)
	searchResponse, err := sarwlerClient.Search(context.Background(), &pbs.SearchRequest{Q: "client good!"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(searchResponse)
}
