package services

import (
	"NewProjectTestingApi/entities"
	"NewProjectTestingApi/repositories"
	"context"
	"gorm.io/gorm"
)

func LoggingIfError(ctx context.Context, params entities.LogbookInsertParams, db *gorm.DB) int {
	return repositories.LoggingIfError(ctx, db, params)
}
