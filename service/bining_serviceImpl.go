package service

import (
	"NewProjectTestingApi/helper"
	"NewProjectTestingApi/model/web"
	"NewProjectTestingApi/repository"
	"context"
	"gorm.io/gorm"
)

type BinningServiceImpl struct {
	DB   *gorm.DB
	Repo repository.BinningRepository
}

func NewBinningServiceImpl(DB *gorm.DB, repo repository.BinningRepository) BinningService {
	return &BinningServiceImpl{DB: DB, Repo: repo}
}

func (service *BinningServiceImpl) FindAll(ctx context.Context, Request []web.BinningHeaderRequest) []web.BinningHeaderResponses {
	//TODO implement me
	allBinning := service.Repo.FindAll(ctx, service.DB, Request)
	return helper.ToHeaderResponses(allBinning)
}
