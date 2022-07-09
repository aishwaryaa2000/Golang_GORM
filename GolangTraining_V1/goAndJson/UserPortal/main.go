package main

import (
	"UserPortal/login"
	"UserPortal/register"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", login.UserLogin).Methods("GET")
	r.HandleFunc("/signup", register.UserRegister).Methods("POST")
	fmt.Println("Server started")
	http.ListenAndServe(":9090", r)
}
