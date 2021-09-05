package services

import (
	"context"
	"grpc-kard/src/services/pbs"
)

type OCRService struct{}

func (obj OCRService) Recognition(ctx context.Context, request *pbs.VideoRequest) (*pbs.VideoResponse, error) {
	return &pbs.VideoResponse{Text: []string{"字幕1 " + string(request.Video.Bt)}}, nil
}
