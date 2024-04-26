package services

import (
	"NewProjectTestingApi/entities"
	payloads2 "NewProjectTestingApi/payloads"
	"context"
)

type BinningService interface {
	FindAll(ctx context.Context, Request []payloads2.BinningHeaderRequest, Log *entities.LogbookInsertParams) ([]payloads2.BinningHeaderResponses, error, string)
}
