package helper

import (
	"NewProjectTestingApi/entities"
	"NewProjectTestingApi/payloads"
	"NewProjectTestingApi/payloads/auth"
	"encoding/json"
	"net/http"
)

func ItemToHeaderResponse(binningDetail entities.BinningStockDetail) payloads.BinningDetailResponses {
	return payloads.BinningDetailResponses{
		BinDocNo:  binningDetail.BinDocNo,
		BinLineNo: binningDetail.BinLineNo,
		PoLineNo:  binningDetail.PoLineNo,
		ItemCode:  binningDetail.ItemCode,
		LocCode:   binningDetail.LocCode,
		CaseNo:    binningDetail.CaseNo,
		GrpoQty:   binningDetail.GrpoQty,
	}
}
func ItemToHeaderResponses(binningDetail []entities.BinningStockDetail) []payloads.BinningDetailResponses {
	var detailResponses []payloads.BinningDetailResponses
	for _, i := range binningDetail {
		detailResponses = append(detailResponses, ItemToHeaderResponse(i))
	}
	return detailResponses
}
func ToHeaderResponse(binningHeader entities.BinningStockHeader) payloads.BinningHeaderResponses {
	return payloads.BinningHeaderResponses{
		CompanyCode: binningHeader.CompanyCode,
		PoDocNo:     binningHeader.PoDocNo,
		WHSGroup:    binningHeader.WHSGroup,
		WHSCode:     binningHeader.WHSCode,
		Item:        ItemToHeaderResponses(binningHeader.Item),
	}

}
func ToHeaderResponses(binningHeader []entities.BinningStockHeader) []payloads.BinningHeaderResponses {
	var binningResponses []payloads.BinningHeaderResponses
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
func ToDomainRegister(userInput auth.RegisterRequest) entities.User {
	return entities.User{
		UserName:   userInput.UserName,
		UserEmail:  userInput.UserEmail,
		Password:   userInput.Password,
		UserRoleId: userInput.UserRoleId,
	}
}
func ToAuthResponses(writer http.ResponseWriter, msg string, status int) auth.AuthResponses {
	writer.WriteHeader(status)
	IsSuccess := true
	if status != 200 { //status berani not http status ok
		IsSuccess = false
	}
	return auth.AuthResponses{
		Message: msg,
		Status:  status,
		Success: IsSuccess,
	}
}
