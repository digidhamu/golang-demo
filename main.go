package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Item Created"))
	fmt.Println("Accessed /create")
}

func HealthCheck(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/create", CreateEndpoint).Methods("GET")
	router.HandleFunc("/health", HealthCheck).Methods("GET")

	fmt.Println("Server is starting on http://:8080/create")

	log.Fatal(http.ListenAndServe(":8080", router))
}
