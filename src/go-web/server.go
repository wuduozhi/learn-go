package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"go-web/controller"
)

func main() {

	mux := httprouter.New()

	mux.ServeFiles("/static/*filepath", http.Dir("public"))
	mux.GET("/hello/:name",controller.Hello)
	mux.GET("/users",controller.GetAllUsers)
	mux.GET("/login",controller.LoginTemplate)
	mux.POST("/authenticate",controller.Login)
	mux.GET("/signup", controller.Signup)
	mux.POST("/signup", controller.DoSignup)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
