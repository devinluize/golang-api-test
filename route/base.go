package route

import (
	"NewProjectTestingApi/app"
	"NewProjectTestingApi/controllers"
	"NewProjectTestingApi/controllers/auth_controller"
	"NewProjectTestingApi/middleware"
	"NewProjectTestingApi/repositories"
	"NewProjectTestingApi/services"
	"github.com/gorilla/mux"
	"net/http"
)

func setupRoute() *mux.Router {
	BinningRepo := repositories.NewBinningRepositoryImpl()
	BinningService := services.NewBinningServiceImpl(app.DB, BinningRepo)
	BinningController := controllers.NewBinningControllerImpl(BinningService)
	//router := httprouter.New()
	routers := mux.NewRouter()
	//router.POST("/api/binning", BinningController.FindAll)

	ProtectedRoute := routers.PathPrefix("").Subrouter()

	ProtectedRoute.HandleFunc("/api/binning", BinningController.FindAll).Methods("POST")

	ProtectedRoute.HandleFunc("/api/Authentication/register/{role}", auth_controller.Register).Methods("POST")
	routers.HandleFunc("/api/Authentication/login", auth_controller.Login).Methods("POST")

	ProtectedRoute.Use(middleware.RouterMiddleware)
	return routers
	//router.POST("/api/authentication/register")

}
func RunRoute() {
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: setupRoute(),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
