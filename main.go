package main

import (
	"NewProjectTestingApi/app"
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
func main() {
	//panic("tes")
	//dbb := app.NewDB()
	app.NewDB()
	route.RunRoute()
}
