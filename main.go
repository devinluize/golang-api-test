package main

import (
	"NewProjectTestingApi/app"
	"NewProjectTestingApi/controllers"
	"NewProjectTestingApi/controllers/auth_controller"
	"NewProjectTestingApi/repositories"
	"NewProjectTestingApi/services"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func MyMiddlewares(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your middleware logic goes here
		fmt.Println("Executing middleware before the handler")
		next.ServeHTTP(w, r)
		fmt.Println("Executing middleware after the handler")
	})
}
func main() {
	//panic("tes")
	//dbb := app.NewDB()
	app.NewDB()
	BinningRepo := repositories.NewBinningRepositoryImpl()
	BinningService := services.NewBinningServiceImpl(app.DB, BinningRepo)
	BinningController := controllers.NewBinningControllerImpl(BinningService)
	//router := httprouter.New()
	routers := mux.NewRouter()
	//router.POST("/api/binning", BinningController.FindAll)
	routers.HandleFunc("/api/binning", BinningController.FindAll).Methods("POST")
	routers.HandleFunc("/api/Authentication/register/{role}", auth_controller.Register).Methods("POST")
	//routers.HandleFunc("/api/Authentication/register/{role}", auth_controller.RegisterHashMicro).Methods("POST")
	routers.HandleFunc("/api/Authentication/login", auth_controller.Login).Methods("POST")

	//router.POST("/api/authentication/register")
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: routers,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
