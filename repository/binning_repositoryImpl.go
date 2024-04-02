package repository

import (
	"NewProjectTestingApi/helper"
	"NewProjectTestingApi/model/domain"
	"NewProjectTestingApi/model/web"
	"context"
	"gorm.io/gorm"
)

type BinningRepositoryImpl struct {
}

func NewBinningRepositoryImpl() BinningRepository {
	return &BinningRepositoryImpl{}
}

func (b *BinningRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB, binning []web.BinningHeaderRequest) []domain.BinningHeader {
	//TODO implement me
	var BinningHeader []domain.BinningHeader
	tx := db.Begin()
	for _, i := range binning {
		var binningStocks domain.BinningHeader
		errs := tx.Raw("EXEC GetHeaderByPODocNoAndCompanyCode @PoDocNo = ?, @CompanyCode = ?", i.PoDocNo, i.CompanyCode).Scan(&binningStocks)
		helper.PanifIfError(errs.Error)
		var BinningStockDetail []domain.BinningStockDetail
		result := tx.Raw("EXEC GetBinningStock @RefDocNo = ?, @CompanyCode = ?", i.PoDocNo, i.CompanyCode).Scan(&BinningStockDetail)
		helper.PanifIfError(result.Error)
		binningStocks.Item = BinningStockDetail
		BinningHeader = append(BinningHeader, binningStocks)
	}

	return BinningHeader
}
