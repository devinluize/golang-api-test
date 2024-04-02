package app

import (
	"NewProjectTestingApi/helper"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() *gorm.DB {
	dsn := "sqlserver://binus_intern:Binus1an@10.1.32.65:1433?database=DMSLIVE_TESTING"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	helper.PanifIfError(err)
	return db
}
