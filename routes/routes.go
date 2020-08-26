package routes

import (
	"github.com/omer-akbas/stock-api/middleware"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	r.HandleFunc("/", indexGet).Methods("GET")
	r.HandleFunc("/api/stock/{code}", stockDetailsGet).Methods("GET")
	r.HandleFunc("/api/stock/{code}/rates", stockRatesGet).Methods("GET")
	r.HandleFunc("/api/stocks/rates", stockListRatesGet).Methods("GET")
	return r
}
