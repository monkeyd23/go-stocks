package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"stocks/api"
)

func main() {
	_ = http.ListenAndServe(port(), setUpRouter())
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func setUpRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", api.Index).Methods("GET")
	r.HandleFunc("/stock/{symbol}", api.GetStocks).Methods("GET")
	return r
}