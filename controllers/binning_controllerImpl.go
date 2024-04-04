package controllers

import (
	"NewProjectTestingApi/helper"
	"NewProjectTestingApi/payloads"
	"NewProjectTestingApi/services"
	"encoding/json"
	"net/http"
)

type BinningControllerImpl struct {
	BinningService services.BinningService
}

func NewBinningControllerImpl(binningService services.BinningService) BinningController {
	return &BinningControllerImpl{BinningService: binningService}
}
func MakeErrorResponse(errMsg string, errorResponses *payloads.ErrorRespones) {
	errorResponses.Success = false
	errorResponses.LogSysNo = 0
	errorResponses.Message = errMsg
}

func (Controller *BinningControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	var binningBodyRequest []payloads.BinningHeaderRequest
	helper.ReadFromRequestBody(request, &binningBodyRequest)
	BinningResponses, errs, errMsg := Controller.BinningService.FindAll(request.Context(), binningBodyRequest, request)

	if errs != nil {
		var errorResponses payloads.ErrorRespones
		MakeErrorResponse(errMsg, &errorResponses)
		err := encoder.Encode(errorResponses)
		helper.PanifIfError(err)
		return
	}
	err := encoder.Encode(BinningResponses)
	helper.PanifIfError(err)

}
