package unittest

import (
	"NewProjectTestingApi/helper"
	"NewProjectTestingApi/model/domain"
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetBinningHeader(t *testing.T) {
	//TODO implement me
	var BinningHeader []domain.BinningHeader
	tx := db.Begin()
	for i := 0; i < 5; i++ {
		var binningStocks domain.BinningHeader
		errs := tx.Raw("EXEC GetHeaderByPODocNoAndCompanyCode @PoDocNo = ?, @CompanyCode = ?", "SPO-NMDI/N/SP/11/20/00183", "3125098").Scan(&binningStocks)
		helper.PanifIfError(errs.Error)
		BinningHeader = append(BinningHeader, binningStocks)
	}
	jsonData, err := json.Marshal(BinningHeader)
	helper.PanifIfError(err)

	fmt.Println(string(jsonData))

}
