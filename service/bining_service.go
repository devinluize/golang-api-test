package service

import (
	"NewProjectTestingApi/model/web"
	"context"
)

type BinningService interface {
	FindAll(ctx context.Context, Request []web.BinningHeaderRequest) []web.BinningHeaderResponses
}
