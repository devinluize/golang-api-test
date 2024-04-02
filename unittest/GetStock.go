package unittest

import (
	"NewProjectTestingApi/model/domain"
	"gorm.io/gorm"
)

func GetBinningStock(db *gorm.DB, refDocNo, companyCode string) ([]domain.BinningStockDetail, error) {
	var binningStocks []domain.BinningStockDetail
	if err := db.Raw("EXEC GetBinningStock @RefDocNo = ?, @CompanyCode = ?", refDocNo, companyCode).Scan(&binningStocks).Error; err != nil {
		return nil, err
	}

	return binningStocks, nil
}
