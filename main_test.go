package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/list", ListEndpoint).Methods("GET")
	router.HandleFunc("/health", HealthEndpoint).Methods("GET")
	return router
}

func TestListEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/list", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestHealthEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/health", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
