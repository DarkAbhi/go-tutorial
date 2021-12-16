package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users/", CreateUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", router))

}

func main() {
	InitialMigration()
	initializeRouter()
}
