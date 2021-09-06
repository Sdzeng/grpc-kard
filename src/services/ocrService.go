package services

import (
	"context"
	"grpc-kard/src/services/pbs"
	"io"
	"log"
	"strconv"
)

type OCRService struct{}

func (obj OCRService) Recognition(ctx context.Context, request *pbs.VideoRequest) (*pbs.VideoResponse, error) {
	return &pbs.VideoResponse{Text: []string{"字幕1 " + string(request.Video.Bt)}}, nil
}

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
