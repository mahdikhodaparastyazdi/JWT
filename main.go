package main

import "github.com/gorilla/mux"
import "jwt/handler"

var router *mux.Router

func CreateRouter() {
	router = mux.NewRouter()
}

func InitializeRoute() {
	router.HandleFunc("/signup", handler.SignUp).Methods("POST")
	router.HandleFunc("/signin", handler.SignIn).Methods("POST")
}

func main() {
	CreateRouter()
	InitializeRoute()
}
