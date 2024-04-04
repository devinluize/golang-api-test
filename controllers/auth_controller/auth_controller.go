package auth_controller

import (
	"NewProjectTestingApi/app"
	"NewProjectTestingApi/helper"
	"NewProjectTestingApi/payloads/auth"
	"NewProjectTestingApi/services/authServices"
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
)

func OpenCon() *gorm.DB {
	dsn := "sqlserver://binus_intern:Binus1an@10.1.32.65:1433?database=DMSLIVE_TESTING"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	//dsn := "sqlserver://binus_intern:Binus1an@10.1.32.65:1433?database=DMSLIVE_TESTING"

	if err != nil {
		panic(err)
	}
	return db
}

var db2 = OpenCon()

func Register(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	role := vars["role"]
	var userInput auth.RegisterRequest
	decoder := json.NewDecoder(request.Body)
	encoder := json.NewEncoder(writer)
	var response auth.RegisterResponses
	writer.Header().Add("Content-Type", "application/json")
	if err := decoder.Decode(&userInput); err != nil {
		response = helper.ToRegisterResponses("Bad Request Format Wrong", http.StatusBadRequest)
		writer.WriteHeader(http.StatusBadRequest)
		err := encoder.Encode(response)
		helper.PanifIfError(err)
		return
	}
	helper.CloseBody(request.Body)
	if role == "admin" {
		userInput.UserRoleId = 1
	} else if role == "hash-micro" {
		userInput.UserRoleId = 2
	} else {
		http.Error(writer, "End Point Wrong", http.StatusBadRequest)
	}
	msg, err := authServices.Register(userInput, app.DB)
	if err != nil {
		response = helper.ToRegisterResponses(msg, http.StatusBadRequest)
		panic(err)
	}
	response = helper.ToRegisterResponses(msg, http.StatusOK)
	err = encoder.Encode(response)
	helper.PanifIfError(err)
}
func Login(writer http.ResponseWriter, request *http.Request) {

}
