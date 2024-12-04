package main

import (
	"log"
	"net/http"
	"my-go-service/middlewares"
	"my-go-service/utils"
)

func main() {
	// Create a new ServeMux instance
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/health", utils.HealthCheckHandler)
	mux.HandleFunc("/metrics", middlewares.MetricsHandler)
	

	handler := middlewares.LoggingMiddleware(
		middlewares.RateLimitMiddleware(
			middlewares.ErrorHandlingMiddleware(
				mux,
			),
		),
	)

	// Start the server
	log.Println("Server starting on port 8080")
	http.ListenAndServe(":8080", handler)
}
