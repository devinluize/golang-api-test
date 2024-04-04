package services

import (
	payloads2 "NewProjectTestingApi/payloads"
	"context"
	"net/http"
)

type BinningService interface {
	FindAll(ctx context.Context, Request []payloads2.BinningHeaderRequest, request *http.Request) ([]payloads2.BinningHeaderResponses, error, string)
}
