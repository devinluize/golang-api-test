package repository

import (
	"NewProjectTestingApi/model/domain"
	"NewProjectTestingApi/model/web"
	"context"
	"gorm.io/gorm"
)

type BinningRepository interface {
	FindAll(ctx context.Context, db *gorm.DB, binning []web.BinningHeaderRequest) []domain.BinningHeader
}
