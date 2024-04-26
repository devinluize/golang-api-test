package route

import (
	"NewProjectTestingApi/app"
	"NewProjectTestingApi/controllers"
	"NewProjectTestingApi/controllers/auth_controller"
	_ "NewProjectTestingApi/docs"
	"NewProjectTestingApi/middleware"
	"NewProjectTestingApi/repositories"
	"NewProjectTestingApi/services"
	"github.com/go-chi/chi"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"net/http"
)

func MyMiddleware(c *fiber.Ctx) error {
	// Your middleware logic here
	return c.Next()
}

func setupRoute() *mux.Router {

	BinningRepo := repositories.NewBinningRepositoryImpl()
	BinningService := services.NewBinningServiceImpl(app.DB, BinningRepo)
	BinningController := controllers.NewBinningControllerImpl(BinningService)
	//router := httprouter.New()
	routers := mux.NewRouter()
	//router.POST("/api/binning", BinningController.FindAll)
	ProtectedRoute := routers.PathPrefix("").Subrouter()
	//routers.HandleFunc("/documentation").Handler(httpSwagger.WrapHandler)
	//ProtectedRoute.HandleFunc("/documentation", httpSwagger.WrapHandler)
	ProtectedRoute.HandleFunc("/swagger/*", httpSwagger.WrapHandler).Methods("get")
	ProtectedRoute.HandleFunc("/api/binning", BinningController.FindAll).Methods("POST")
	ProtectedRoute.HandleFunc("/api/Authentication/register/{role}", auth_controller.Register).Methods("POST")
	routers.HandleFunc("/api/Authentication/login", auth_controller.Login).Methods("POST")
	ProtectedRoute.Use(middleware.RouterMiddleware)

	//////////////////////

	//fib := fiber.New()
	//fib.Get("/swagger/*", swagger.HandlerDefault)
	//fib.Post("/api/binning")
	//fib.Post("/api/binning", MyMiddleware, BinningController.FindAll)
	return routers
	//router.POST("/api/authentication/register")

}

func setupRouteUsingCHi() *chi.Mux {

	BinningRepo := repositories.NewBinningRepositoryImpl()
	BinningService := services.NewBinningServiceImpl(app.DB, BinningRepo)
	BinningController := controllers.NewBinningControllerImpl(BinningService)
	routers3 := chi.NewRouter()
	routers3.Post("/api/binning", BinningController.FindAll)
	routers3.With(middleware.RouterMiddleware).Post("/api/Authentication/register/{role}", BinningController.FindAll)
	routers3.Post("/api/Authentication/login", auth_controller.Login)
	routers3.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3030/swagger/doc.json"),
	))
	return routers3
}

func RunRoute() {
	r := chi.NewRouter()
	//r.Get("/swagger/*", httpSwagger.Handler())
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	//r.Get("/swagger/*", httpSwagger.Handler(
	//	httpSwagger.URL("http://localhost:1323/swagger/docs/swagger.json"), //The url pointing to API definition
	//))

	server := http.Server{
		Addr:    "localhost:3030",
		Handler: setupRouteUsingCHi(),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
