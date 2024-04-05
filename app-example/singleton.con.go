package app_example

import (
	"NewProjectTestingApi/helper"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func NewDB() {
	dsn := "sqlserver://username:password@ip:port?database=databasename"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	helper.PanifIfError(err)
	DB = db
}
