package controllers

import (
	"NewProjectTestingApi/entities"
	"NewProjectTestingApi/services"
	"context"
	"gorm.io/gorm"
)

func LoggingIFError(ctx context.Context, db *gorm.DB, params entities.LogbookInsertParams) int {
	return services.LoggingIfError(ctx, params, db)
}
