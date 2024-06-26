package repositories

import (
	"NewProjectTestingApi/entities"
	"NewProjectTestingApi/payloads"
	"context"
	"gorm.io/gorm"
)

type BinningRepository interface {
	FindAll(ctx context.Context, db *gorm.DB, binning []payloads.BinningHeaderRequest, Log *entities.LogbookInsertParams) ([]entities.BinningStockHeader, error, string)
}
