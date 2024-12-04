package main

import (
	"log"
	"net/http"
	"my-go-service/middlewares"
	"my-go-service/utils"

	"github.com/gorilla/mux"
)

func main() {

	mux := http.NewServeMux()

	// Middlewares
	ux.HandleFunc(middlewares.LoggingMiddleware)
	ux.HandleFunc(middlewares.RateLimitMiddleware)
	ux.HandleFunc(middlewares.ErrorHandlingMiddleware)
	ux.HandleFunc(middlewares.AuthMiddleware)

	// Routes
	router.HandleFunc("/health", utils.HealthCheckHandler).Methods("GET")
	router.Handle("/metrics", middlewares.MetricsHandler())

	log.Println("Server starting on port 8080")
	http.ListenAndServe(":8080", router)
}
