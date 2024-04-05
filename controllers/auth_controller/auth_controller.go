package auth_controller

import (
	"NewProjectTestingApi/app"
	"NewProjectTestingApi/config"
	"NewProjectTestingApi/helper"
	"NewProjectTestingApi/payloads/auth"
	"NewProjectTestingApi/services/authServices"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func Register(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	role := vars["role"]
	var userInput auth.RegisterRequest
	decoder := json.NewDecoder(request.Body)
	encoder := json.NewEncoder(writer)
	var response auth.AuthResponses
	writer.Header().Add("Content-Type", "application/json")
	if err := decoder.Decode(&userInput); err != nil {
		response = helper.ToAuthResponses(writer, "Bad Request Format Wrong", http.StatusBadRequest)
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
		writer.WriteHeader(http.StatusBadRequest)
		err := encoder.Encode(response)
		helper.PanifIfError(err)
	}
	msg, err := authServices.Register(userInput, app.DB)
	if err != nil {
		response = helper.ToAuthResponses(writer, msg, http.StatusBadRequest)
		writer.WriteHeader(http.StatusBadRequest)
		err := encoder.Encode(response)
		helper.PanifIfError(err)
	}
	response = helper.ToAuthResponses(writer, msg, http.StatusOK)
	writer.WriteHeader(http.StatusBadRequest)
	err = encoder.Encode(response)
	helper.PanifIfError(err)
}
func Login(writer http.ResponseWriter, request *http.Request) {
	var userInput auth.LoginRequest
	decoder := json.NewDecoder(request.Body)
	encoder := json.NewEncoder(writer)
	var response auth.AuthResponses
	writer.Header().Add("Content-Type", "application/json")
	if err := decoder.Decode(&userInput); err != nil {
		response = helper.ToAuthResponses(writer, "Format Login Wrong", http.StatusBadRequest)
		return
	}
	helper.CloseBody(request.Body)
	msg, err, roleid := authServices.Login(userInput, app.DB)
	if err != nil {
		response = helper.ToAuthResponses(writer, msg, http.StatusUnauthorized)
		writer.WriteHeader(http.StatusBadRequest)
		err := encoder.Encode(response)
		helper.PanifIfError(err)
		return
	}
	expTime := time.Now().Add(time.Minute * 10)
	claims := config.JWTClaim{
		UserName: userInput.UserName,
		UserRole: roleid,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "devin",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response = helper.ToAuthResponses(writer, err.Error(), http.StatusBadRequest)
		writer.WriteHeader(http.StatusBadRequest)
		err := encoder.Encode(response)
		helper.PanifIfError(err)
		return
	}
	http.SetCookie(
		writer, &http.Cookie{
			Name:     "token",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
		})
	response = helper.ToAuthResponses(writer, "Success Login", http.StatusOK)
	writer.WriteHeader(http.StatusOK)
	err = encoder.Encode(response)
	helper.PanifIfError(err)
	return
}
