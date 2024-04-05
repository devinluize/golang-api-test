package config

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("dasdasdasdas")

type JWTClaim struct {
	UserName string
	UserRole int
	jwt.RegisteredClaims
}
