package repositories

import (
	"NewProjectTestingApi/entities"
	"NewProjectTestingApi/payloads"
	"context"
	"gorm.io/gorm"
	"net/http"
)

type BinningRepository interface {
	FindAll(ctx context.Context, db *gorm.DB, binning []payloads.BinningHeaderRequest, httprequest *http.Request) ([]entities.BinningStockHeader, error, string)
}
