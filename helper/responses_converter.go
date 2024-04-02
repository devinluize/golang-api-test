package helper

import (
	"NewProjectTestingApi/model/domain"
	"NewProjectTestingApi/model/web"
	"encoding/json"
	"net/http"
)

func ItemToHeaderResponse(binningDetail domain.BinningStockDetail) web.BinningDetailResponses {
	return web.BinningDetailResponses{
		BinDocNo:  binningDetail.BinDocNo,
		BinLineNo: binningDetail.BinLineNo,
		PoLineNo:  binningDetail.PoLineNo,
		ItemCode:  binningDetail.ItemCode,
		LocCode:   binningDetail.LocCode,
		CaseNo:    binningDetail.CaseNo,
		GrpoQty:   binningDetail.GrpoQty,
	}
}
func ItemToHeaderResponses(binningDetail []domain.BinningStockDetail) []web.BinningDetailResponses {
	var detailResponses []web.BinningDetailResponses
	for _, i := range binningDetail {
		detailResponses = append(detailResponses, ItemToHeaderResponse(i))
	}
	return detailResponses
}
func ToHeaderResponse(binningHeader domain.BinningHeader) web.BinningHeaderResponses {
	return web.BinningHeaderResponses{
		CompanyCode: binningHeader.CompanyCode,
		PoDocNo:     binningHeader.PoDocNo,
		WHSGroup:    binningHeader.WHSGroup,
		WHSCode:     binningHeader.WHSCode,
		Item:        ItemToHeaderResponses(binningHeader.Item),
	}

}
func ToHeaderResponses(binningHeader []domain.BinningHeader) []web.BinningHeaderResponses {
	var binningResponses []web.BinningHeaderResponses
	for _, i := range binningHeader {

		binningResponses = append(binningResponses, ToHeaderResponse(i))
	}
	return binningResponses
}
func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	if err != nil {
		panic(err)
	}
}
