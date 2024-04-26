package main

import (
	"NewProjectTestingApi/app"
	_ "NewProjectTestingApi/docs"
	"NewProjectTestingApi/route"
	"fmt"
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

//func tes2(Log *namee) {
//	Log.name = "dadasdas"
//}
//func teseditdata(Log *namee) {
//	Log.name = "dev"
//	tes2(Log)
//}
//
//type namee struct {
//	name string
//}

// @title			Binning List
// @version		1.0
// @description	This is a practice api spec for getting binning list.
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host
// @BasePath	/http://localhost:3030/
func main() {
	//panic("tes")
	//dbb := app.NewDB()

	app.NewDB()
	//route.RunRoute()

	//mux router
	//r := mux.NewRouter()
	//
	//r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
	//	httpSwagger.URL("http://localhost:3030/swagger/doc.json"), //The url pointing to API definition
	//	httpSwagger.DeepLinking(true),
	//	httpSwagger.DocExpansion("none"),
	//	httpSwagger.DomID("swagger-ui"),
	//)).Methods(http.MethodGet)
	//chi

	//r := chi.NewRouter()
	////r.Mount("/swagger", httpSwagger.WrapHandler)
	//r.Get("/swagger/*", httpSwagger.Handler(
	//	httpSwagger.URL("http://localhost:3030/swagger/doc.json"), // URL to your Swagger JSON endpoint
	//))

	//r.Get("/swagger/*", httpSwagger.Handler())
	//r.Get("/swagger/*", httpSwagger.WrapHandler)

	//r.Get("/swagger/*", httpSwagger.Handler(
	//	httpSwagger.URL("http://localhost:1323/swagger/docs/swagger.json"), //The url pointing to API definition
	//))

	//server := http.Server{
	//	Addr:    "localhost:3030",
	//	Handler: r,
	//}
	//server.ListenAndServe()
	route.RunRoute()
	//name := namee{name: "yeb"}
	//teseditdata(&name)
	//fmt.Println(name.name)

}
