package main

import (
	"encoding/json"
	"log"
	"net/http"
	"my-go-service/middlewares"
	"my-go-service/utils"
	"sync"
)
var (
	// Simulated user database
	users = map[string]string{} // username:password
	mu    sync.Mutex            // mutex to protect the user map
)

// SignupHandler handles user signup
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Add user to the database
	mu.Lock()
	defer mu.Unlock()
	if _, exists := users[creds.Username]; exists {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}
	users[creds.Username] = creds.Password
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

// LoginHandler handles user login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Check user credentials
	mu.Lock()
	defer mu.Unlock()
	if password, exists := users[creds.Username]; !exists || password != creds.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}
func main() {
	// Create a new ServeMux instance
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/health", utils.HealthCheckHandler)
	mux.HandleFunc("/metrics", middlewares.MetricsHandler)
	mux.HandleFunc("/signup", SignupHandler)
	mux.HandleFunc("/login", LoginHandler)

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
