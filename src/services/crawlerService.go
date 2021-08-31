package services

import (
	"context"
	"encoding/json"
	"grpc-kard/src/services/pbs"
)

type CrawlerService struct{}

func (obj CrawlerService) Search(cxt context.Context, request *pbs.SearchRequest) (*pbs.SearchResponse, error) {
	return &pbs.SearchResponse{Text: "q=" + request.Q}, nil
}

func (obj CrawlerService) New(cxt context.Context, request *pbs.NewRequest) (*pbs.NewResponse, error) {
	bytes, _ := json.Marshal(request.Nm)
	return &pbs.NewResponse{Text: string(bytes)}, nil
}
