package services

import (
	"NewProjectTestingApi/helper"
	payloads2 "NewProjectTestingApi/payloads"
	"NewProjectTestingApi/repositories"
	"context"
	"gorm.io/gorm"
	"net/http"
)

type BinningServiceImpl struct {
	DB   *gorm.DB
	Repo repositories.BinningRepository
}

func NewBinningServiceImpl(DB *gorm.DB, repo repositories.BinningRepository) BinningService {
	return &BinningServiceImpl{DB: DB, Repo: repo}
}

func (service *BinningServiceImpl) FindAll(ctx context.Context, Request []payloads2.BinningHeaderRequest, httprequest *http.Request) ([]payloads2.BinningHeaderResponses, error, string) {
	//TODO implement me
	allBinning, err, errMsg := service.Repo.FindAll(ctx, service.DB, Request, httprequest)
	return helper.ToHeaderResponses(allBinning), err, errMsg
}
