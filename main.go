package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ListEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Items List"))
	fmt.Println("Accessed /list")
}

func HealthEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/list", ListEndpoint).Methods("GET")
	router.HandleFunc("/health", HealthEndpoint).Methods("GET")

	fmt.Println("Server is starting on http://:8080/list")

	log.Fatal(http.ListenAndServe(":8080", router))
}
