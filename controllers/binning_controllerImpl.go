package controllers

import (
	"NewProjectTestingApi/entities"
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

// FindAll Get All Bining List Via Header
//
//	@Summary		Show An Binning List
//	@Description	Get Binning List By Header
//	@Tags			Tags1Test
//	@Accept			json
//	@Produce		json
//	@Param			request	body		[]payloads.BinningHeaderRequest	true	"Insert Header Request"
//	@Success		200		{object}	[]payloads.BinningHeaderResponses
//	@Router			/api/binning [post]
func (Controller *BinningControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	v := entities.LogbookInsertParams{}
	headerString, err2 := json.Marshal(request.Header)
	helper.PanifIfError(err2)
	v.RequestHeader = string(headerString)
	encoder := json.NewEncoder(writer)
	var binningBodyRequest []payloads.BinningHeaderRequest
	helper.ReadFromRequestBody(request, &binningBodyRequest)
	BinningResponses, errs, errMsg := Controller.BinningService.FindAll(request.Context(), binningBodyRequest, &v)

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
