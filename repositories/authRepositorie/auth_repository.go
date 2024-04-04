package authRepositorie

import (
	"NewProjectTestingApi/helper"
	"NewProjectTestingApi/payloads/auth"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(userInput auth.RegisterRequest, db *gorm.DB) (string, error) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)
	user := helper.ToDomainRegister(userInput)
	err := db.Create(user)
	if err.Error != nil {

		return "Register To Database Gagal Bad Responses", err.Error
	}
	return "Register Success", nil
}
