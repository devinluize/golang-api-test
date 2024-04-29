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
	username := "binus_intern"
	password := "Binus1an"
	port := "1433" // default port for SQL Server is 1433
	database := "DMSLIVE_TESTING"
	serverIP := "10.1.32.65"
	// Create the DSN connection string
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", username, password, serverIP, port, database)
	//dsn := "sqlserver://binus_intern:Binus1an@10.1.32.65:1433?database=DMSLIVE_TESTING"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	//dsn := "sqlserver://binus_intern:Binus1an@10.1.32.65:1433?database=DMSLIVE_TESTING"
	helper.PanifIfError(err)
	//return db
	DB = db
}
