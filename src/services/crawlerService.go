package services

import (
	"context"
	"grpc-kard/src/services/pbs"
)

type CrawlerService struct{}

func (obj CrawlerService) Search(context.Context, *pbs.SearchRequest) (*pbs.SearchResponse, error) {
	return &pbs.SearchResponse{Text: "good"}, nil
}
