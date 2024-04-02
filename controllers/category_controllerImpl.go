package controllers

import (
	"NewProjectTestingApi/helper"
	"NewProjectTestingApi/payloads"
	"NewProjectTestingApi/services"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type BinningControllerImpl struct {
	BinningService services.BinningService
}

func NewBinningControllerImpl(binningService services.BinningService) BinningController {
	return &BinningControllerImpl{BinningService: binningService}
}

func (Controller *BinningControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	var binningBodyRequest []payloads.BinningHeaderRequest
	helper.ReadFromRequestBody(request, &binningBodyRequest)
	BinningResponses := Controller.BinningService.FindAll(request.Context(), binningBodyRequest)
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(BinningResponses)
	helper.PanifIfError(err)

}
