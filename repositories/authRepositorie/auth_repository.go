package authRepositorie

import (
	"NewProjectTestingApi/entities"
	"NewProjectTestingApi/helper"
	"NewProjectTestingApi/payloads/auth"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RegisterUser(userInput auth.RegisterRequest, db *gorm.DB) (string, error) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	tx := db.Begin()
	defer tx.Rollback()
	userInput.Password = string(hashPassword)
	user := helper.ToDomainRegister(userInput)
	err := tx.Create(user)
	if err == nil {
		tx.Commit()
	}
	if err.Error != nil {

		return "Register To Database Gagal Bad Responses", err.Error
	}
	return "Register Success", nil
}
func LoginUser(userInput auth.LoginRequest, db *gorm.DB) (string, error, int) {
	var user entities.User
	if err := db.Where("user_name =?", userInput.UserName).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return "username atau password salah", err, -1
		default:
			return "Bad Request", err, -1
		}
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
	if err != nil {
		return "Password Salah", err, -1
	}
	return "success", nil, user.UserRoleId

}
