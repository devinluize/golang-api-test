package app

import (
	"NewProjectTestingApi/helper"
	"fmt"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func NewDB() {
	username := "database_username"
	password := "database_password"
	port := "1433"
	database := "database name"
	serverIP := "ip database"
	// Create the DSN connection string
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", username, password, serverIP, port, database)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	helper.PanifIfError(err)
	DB = db
}
