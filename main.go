package main

import (
	"NewProjectTestingApi/app"
	"NewProjectTestingApi/controllers"
	"NewProjectTestingApi/repositories"
	"NewProjectTestingApi/services"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	//panic("tes")
	db := app.NewDB()
	BinningRepo := repositories.NewBinningRepositoryImpl()
	BinningService := services.NewBinningServiceImpl(db, BinningRepo)
	BinningController := controllers.NewBinningControllerImpl(BinningService)
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
