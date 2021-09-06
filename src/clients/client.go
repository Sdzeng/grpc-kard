package main

import (
	"context"
	"fmt"
	"grpc-kard/src/services/pbs"
	"io"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
)

func main() {
	opt := []grpc.DialOption{grpc.WithInsecure(), grpc.WithReturnConnectionError()}
	conn, err := grpc.Dial(":8081", opt...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//clients
	sarwlerClient := pbs.NewCarwlerServiceClient(conn)
	ocrClient := pbs.NewOCRServiceClient(conn)

	ctx := context.Background()
	searchResponse, err := sarwlerClient.Search(ctx, &pbs.SearchRequest{Q: "client good!"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("searchResponse=", searchResponse)

	t := &timestamp.Timestamp{Seconds: time.Now().Unix()}
	newResponse, err := sarwlerClient.New(ctx, &pbs.NewRequest{Nm: &pbs.NewMessage{Id: 1, Time: t}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("newResponse=", newResponse)

	videoResponse, err := ocrClient.Recognition(ctx, &pbs.VideoRequest{Video: &pbs.VideoBytes{Bt: []byte("good bytes")}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("videoResponse=", videoResponse)

	//server stream
	stream, err := ocrClient.RecognitionByServerStream(ctx, &pbs.VideoRequest{Video: &pbs.VideoBytes{Bt: []byte("good bytes")}})
	if err != nil {
		log.Fatal(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(res.Text)
	}

	//client stream
	stream2, err := ocrClient.RecognitionByClientStream(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = stream2.Send(&pbs.VideoRequest{Video: &pbs.VideoBytes{Bt: []byte("good bytes")}})
	if err != nil {
		log.Fatal(err)
	}

	res2, err := stream2.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(res2)
	}

}
