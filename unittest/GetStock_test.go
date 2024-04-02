package unittest

import (
	"NewProjectTestingApi/helper"
	"NewProjectTestingApi/model/domain"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"testing"
)

func OpenCon() *gorm.DB {
	dsn := "sqlserver://binus_intern:Binus1an@10.1.32.65:1433?database=DMSLIVE_TESTING"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	helper.PanifIfError(err)
	return db
}

var db = OpenCon()

func TestGetStock(t *testing.T) {
	var binningStocks []domain.BinningStockDetail
	err := db.Raw("EXEC GetBinningStock @RefDocNo = ?, @CompanyCode = ?", "SPO-NMDI/N/SP/11/20/00183", "3125098").Scan(&binningStocks).Error
	//err := db.Raw("EXEC GetBinningStock @RefDocNo = ?, @CompanyCode = ?", refDocNo, companyCode).Scan(&binningStocks).Error
	for _, bs := range binningStocks {
		fmt.Printf("BinDocNo: %s, BinLineNo: %s, PoLineNo: %s, ItemCode: %s, LocCode: %s, CaseNo: %s, GrpoQty: %d\n",
			bs.BinDocNo, bs.BinLineNo, bs.PoLineNo, bs.ItemCode, bs.LocCode, bs.CaseNo, bs.GrpoQty)
	}
	helper.PanifIfError(err)
	assert.Nil(t, err)
}
func TestGetBinningUsingTrans(t *testing.T) {
	var BinningStockDetail []domain.BinningStockDetail
	result := db.Raw("EXEC GetBinningStock @RefDocNo = ?, @CompanyCode = ?", "SPO-NMDI/N/SP/11/20/00183", "3125098").Scan(&BinningStockDetail)
	helper.PanifIfError(result.Error)
}
func GetBinningStocsks() []domain.BinningStockDetail {
	var binningStocks []domain.BinningStockDetail
	result := db.Raw("EXEC GetBinningStock @RefDocNo = ?, @CompanyCode = ?", "SPO-NMDI/N/SP/11/20/00183", "3125098").Scan(&binningStocks)
	helper.PanifIfError(result.Error)
	return binningStocks

	//return result[1]
}

func TestGetHeader(t *testing.T) {
	var binningStocks domain.BinningStockHeader
	tx := db.Begin()
	// Call the stored procedure GetHeaderByPODocNoAndCompanyCode
	errs := tx.Raw("EXEC GetHeaderByPODocNoAndCompanyCode @PoDocNo = ?, @CompanyCode = ?", "SPO-NMDI/N/SP/11/20/00183", "3125098").Scan(&binningStocks)
	if errs != nil {
		tx.Rollback()
	}

	fmt.Println("Headers:")
	binningStocks.Item = GetBinningStocsks()
	fmt.Println(binningStocks.Item[1])
}
