package repositories

import (
	"NewProjectTestingApi/entities"
	"NewProjectTestingApi/helper"
	"NewProjectTestingApi/payloads"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type BinningRepositoryImpl struct {
}

func NewBinningRepositoryImpl() BinningRepository {
	return &BinningRepositoryImpl{}
}
func GetStockDetail(db *gorm.DB, BinningStockDetail *[]entities.BinningStockDetail, binning payloads.BinningHeaderRequest) error {

	errs := db.Table("atbinningstock0 A").
		Select("A.BIN_DOC_NO, B.BIN_LINE_NO, B.REF_LINE AS PO_LINE_NO, B.ITEM_CODE, B.LOC_CODE, A.SUPPLIER_CASE_NO AS CASE_NO, CAST(B.DO_QTY AS FLOAT) AS GRPO_QTY").
		Joins("INNER JOIN atbinningstock1 B ON A.BIN_SYS_NO = B.BIN_SYS_NO").
		Where("REF_DOC_NO = ? AND COMPANY_CODE = ?", binning.PoDocNo, binning.CompanyCode).
		Find(&BinningStockDetail)
	if errors.Is(errs.Error, gorm.ErrRecordNotFound) {
		recover()
		return nil
	}
	return errs.Error
}
func GetBinningHeader(db *gorm.DB, binning payloads.BinningHeaderRequest, binningStocks *entities.BinningStockHeader) error {
	errs := db.Where("VIA_BINNING = ? AND PO_DOC_NO = ? AND COMPANY_CODE = ?", "1", binning.PoDocNo, binning.CompanyCode).First(&binningStocks)
	//if errors.Is(errs.Error, gorm.ErrRecordNotFound) {
	//	recover()
	//	binningStocks = nil
	//	return nil
	//}
	return errs.Error
}
func Equals(bs entities.BinningStockHeader, other entities.BinningStockHeader) bool {
	return bs.CompanyCode == other.CompanyCode &&
		bs.PoDocNo == other.PoDocNo &&
		bs.WHSGroup == other.WHSGroup &&
		bs.WHSCode == other.WHSCode
}

var log entities.LogbookInsertParams

func populteLogBookParams(r http.Request) entities.LogbookInsertParams {
	//belum selesai develop
	jsonBytes, err := json.Marshal(r)
	helper.PanifIfError(err)
	return entities.LogbookInsertParams{
		ApiName:        "api/binning",
		ApiMethod:      "POST",
		ApiUrl:         "api/binning",
		ApiParameters:  "",
		ApiStart:       time.Now(),
		ApiFinish:      time.Time{},
		ApiStatus:      "404",
		RequestHeader:  r.Header.Get("Content-Type") + r.Header.Get(""),
		RequestBody:    string(jsonBytes),
		ResponseHeader: "",
		ResponseBody:   "",
		RowsAffected:   0,
		Success:        false,
		Message:        "",
		ProcessFrom:    time.Time{},
		ProcessBy:      "",
	}
}

// hashmicro
func (b *BinningRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB, binning []payloads.BinningHeaderRequest, Log *entities.LogbookInsertParams) ([]entities.BinningStockHeader, error, string) {
	//TODO implement me
	Log.ApiName = "api/bininglist"
	Log.ApiMethod = "POST"
	Log.ApiUrl = "api/binning"
	Log.ApiStart = time.Now()

	tx := db.Begin()
	defer tx.Rollback()
	var BinningHeader []entities.BinningStockHeader
	var err error = nil
	var errOnDetail error = nil

	for _, i := range binning {
		var binningStocks entities.BinningStockHeader
		var BinningStockDetail []entities.BinningStockDetail
		err = GetBinningHeader(db, i, &binningStocks)

		if err != nil {
			break
		}
		errOnDetail = GetStockDetail(db, &BinningStockDetail, i)

		if Equals(binningStocks, entities.BinningStockHeader{}) {
			continue
		}
		if err != nil || errOnDetail != nil {
			tx.Rollback()
			fmt.Println("Error Test For Log")
			continue
		}
		binningStocks.Item = BinningStockDetail
		BinningHeader = append(BinningHeader, binningStocks)
	}
	if err != nil || errOnDetail != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return BinningHeader, err, "Record Not Found"
		}
		return BinningHeader, err, "Error Reading to Database"
	}
	return BinningHeader, nil, ""
}
