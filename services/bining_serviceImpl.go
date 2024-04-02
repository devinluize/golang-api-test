package services

import (
	"NewProjectTestingApi/helper"
	payloads2 "NewProjectTestingApi/payloads"
	"NewProjectTestingApi/repositories"
	"context"
	"gorm.io/gorm"
)

type BinningServiceImpl struct {
	DB   *gorm.DB
	Repo repositories.BinningRepository
}

func NewBinningServiceImpl(DB *gorm.DB, repo repositories.BinningRepository) BinningService {
	return &BinningServiceImpl{DB: DB, Repo: repo}
}

func (service *BinningServiceImpl) FindAll(ctx context.Context, Request []payloads2.BinningHeaderRequest) []payloads2.BinningHeaderResponses {
	//TODO implement me
	allBinning := service.Repo.FindAll(ctx, service.DB, Request)
	return helper.ToHeaderResponses(allBinning)
}
