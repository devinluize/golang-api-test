package middleware

import (
	"NewProjectTestingApi/config"
	"NewProjectTestingApi/helper"
	"NewProjectTestingApi/payloads/auth"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func RouterMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("token")
		encoder := json.NewEncoder(writer)
		var response auth.AuthResponses
		writer.Header().Add("Content-Type", "application/json")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				response = helper.ToAuthResponses(writer, "You are Unauthorized please login", http.StatusUnauthorized)
				err := encoder.Encode(response)
				helper.PanifIfError(err)
				return
			}
			response = helper.ToAuthResponses(writer, "Internal server error", http.StatusInternalServerError)
			err := encoder.Encode(response)
			helper.PanifIfError(err)
			return
		}
		tokenString := cookie.Value
		claims := config.JWTClaim{}

		myToken, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})
		helper.PanifIfError(err)
		if claims.UserRole != 1 {
			response = helper.ToAuthResponses(writer, "You are not admin", http.StatusUnauthorized)
			err := encoder.Encode(response)
			helper.PanifIfError(err)
			return
		}
		if !myToken.Valid {
			response = helper.ToAuthResponses(writer, "Token Is Not Valid", http.StatusUnauthorized)
			err := encoder.Encode(response)
			helper.PanifIfError(err)
			return
		}
		handler.ServeHTTP(writer, request)
	})
}
