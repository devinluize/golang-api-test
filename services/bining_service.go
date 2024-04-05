package services

import (
	payloads2 "NewProjectTestingApi/payloads"
	"context"
)

type BinningService interface {
	FindAll(ctx context.Context, Request []payloads2.BinningHeaderRequest) ([]payloads2.BinningHeaderResponses, error, string)
}
