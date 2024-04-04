package controllers

import (
	"net/http"
)

type BinningController interface {
	FindAll(writer http.ResponseWriter, request *http.Request)
}
