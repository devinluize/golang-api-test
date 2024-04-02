package unittest

import (
	"NewProjectTestingApi/entities"
	"gorm.io/gorm"
)

func GetBinningStock(db *gorm.DB, refDocNo, companyCode string) ([]entities.BinningStockDetail, error) {
	var binningStocks []entities.BinningStockDetail
	if err := db.Raw("EXEC GetBinningStock @RefDocNo = ?, @CompanyCode = ?", refDocNo, companyCode).Scan(&binningStocks).Error; err != nil {
		return nil, err
	}

	return binningStocks, nil
}
