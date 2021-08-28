package services

import (
	"context"
	"grpc-kard/src/services/pbs"
)

type CrawlerService struct{}

func (this CrawlerService) Search(context.Context, *pbs.SearchRequest) (*pbs.SearchResponse, error) {
	return &pbs.SearchResponse{Text: ""}, nil
}
