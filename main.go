package main

import (
	"NewProjectTestingApi/app"
	"NewProjectTestingApi/controller"
	"NewProjectTestingApi/repository"
	"NewProjectTestingApi/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	//panic("tes")
	db := app.NewDB()
	BinningRepo := repository.NewBinningRepositoryImpl()
	BinningService := service.NewBinningServiceImpl(db, BinningRepo)
	BinningController := controller.NewBinningControllerImpl(BinningService)
	router := httprouter.New()
	router.POST("/api/binning", BinningController.FindAll)
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
