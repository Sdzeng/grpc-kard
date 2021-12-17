package services

import (
	"context"
	"grpc-kard/src/services/pbs"
	"io"
	"log"
	"strconv"
)

type OCRService struct{}

//普通方式
func (obj OCRService) Recognition(ctx context.Context, request *pbs.VideoRequest) (*pbs.VideoResponse, error) {
	return &pbs.VideoResponse{Text: []string{"字幕1 " + string(request.Video.Bt)}}, nil
}

//服务端流模式
func (obj OCRService) RecognitionByServerStream(request *pbs.VideoRequest, stream pbs.OCRService_RecognitionByServerStreamServer) error {

	textArr := make([]string, 0)
	for i := 1; i < 100; i++ {
		textArr = append(textArr, "good stream coming "+strconv.Itoa(i))
		if i%5 == 0 {
			err := stream.Send(&pbs.VideoResponse{Text: textArr})
			if err != nil {
				log.Fatal(err)
				break
			}
			textArr = (textArr)[0:0]
		}
	}

	if len(textArr) > 0 {
		err := stream.Send(&pbs.VideoResponse{Text: textArr})
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

//客户端流模式
func (obj OCRService) RecognitionByClientStream(stream pbs.OCRService_RecognitionByClientStreamServer) error {

	textArr := make([]string, 0)
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pbs.VideoResponse{Text: textArr})
		}
		if err != nil {
			log.Fatal(err)
		}

		textArr = append(textArr, "text")
	}

}

//双向流模式
func (obj OCRService) RecognitionByTWStream(stream pbs.OCRService_RecognitionByTWStreamServer) error {
	i := 1
	for {
		_, err := stream.Recv()
		if err == io.EOF {

			return nil
		}
		if err != nil {
			log.Fatal(err)
		}

		textArr := []string{"text " + strconv.Itoa(i)}
		i++

		err = stream.Send(&pbs.VideoResponse{Text: textArr})
		if err != nil {
			log.Println(err)
		}
	}
}
