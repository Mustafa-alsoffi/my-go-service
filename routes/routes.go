package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Define response structure for errors.
type ErrorResponse struct {
	Error string `json:"error"`
}

// Handlers for different routes.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the Home Page"))
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This is the About Page"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(ErrorResponse{Error: "Resource not found"})
}

// InitializeRoutes sets up all routes and returns a router.
func InitializeRoutes() *mux.Router {
	router := mux.NewRouter()

	// Define application routes.
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/about", AboutHandler).Methods("GET")

	// Custom 404 handler.
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	return router
}
