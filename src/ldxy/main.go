package main

import (
	"ldxy/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/beat", api.BeatHandler)
	http.HandleFunc("/list", api.ListHandler)
	http.HandleFunc("/signup", api.Signup)
	http.HandleFunc("/manage", api.AdminHandler)
	http.HandleFunc("/tv", api.TvHandler)
	http.HandleFunc("/login", api.LoginHandler)
	//static file handler.
	http.Handle("/", http.FileServer(http.Dir("./static")))
	//Listen on port 8080
	log.Fatal(http.ListenAndServe(":415", nil))
}
