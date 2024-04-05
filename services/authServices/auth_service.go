package authServices

import (
	"NewProjectTestingApi/payloads/auth"
	"NewProjectTestingApi/repositories/authRepositorie"
	"gorm.io/gorm"
)

func Register(userInput auth.RegisterRequest, db *gorm.DB) (string, error) {
	msg, err := authRepositorie.RegisterUser(userInput, db)
	return msg, err
}
func Login(userInput auth.LoginRequest, db *gorm.DB) (string, error, int) {
	msg, err, role := authRepositorie.LoginUser(userInput, db)
	return msg, err, role
}
