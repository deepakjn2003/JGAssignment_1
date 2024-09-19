package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func initRouter(){
	r := mux.NewRouter()

	r.HandleFunc("/", handler_home)//.Methods("GET")
				//"/view/ vid123"
	r.HandleFunc("/view/{vid}", handler_view)//.Methods("POST") {vid : vid123}
	r.HandleFunc("/count/{vid}", handler_viewcount)//.Methods("GET")

	http.ListenAndServe(":8080", r)
}