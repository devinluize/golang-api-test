package repositories

import (
	"NewProjectTestingApi/entities"
	"NewProjectTestingApi/helper"
	"NewProjectTestingApi/payloads"
	"context"
	"gorm.io/gorm"
)

type BinningRepositoryImpl struct {
}

func NewBinningRepositoryImpl() BinningRepository {
	return &BinningRepositoryImpl{}
}

func (b *BinningRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB, binning []payloads.BinningHeaderRequest) []entities.BinningHeader {
	//TODO implement me
	var BinningHeader []entities.BinningHeader
	tx := db.Begin()
	for _, i := range binning {
		var binningStocks entities.BinningHeader
		errs := tx.Raw("EXEC GetHeaderByPODocNoAndCompanyCode @PoDocNo = ?, @CompanyCode = ?", i.PoDocNo, i.CompanyCode).Scan(&binningStocks)
		helper.PanifIfError(errs.Error)
		var BinningStockDetail []entities.BinningStockDetail
		result := tx.Raw("EXEC GetBinningStock @RefDocNo = ?, @CompanyCode = ?", i.PoDocNo, i.CompanyCode).Scan(&BinningStockDetail)
		helper.PanifIfError(result.Error)
		binningStocks.Item = BinningStockDetail
		BinningHeader = append(BinningHeader, binningStocks)
	}

	return BinningHeader
}
