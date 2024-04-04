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
