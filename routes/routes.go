package routes

import (
	"banking-api/internal/handler"
	"net/http"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/health", handler.HealthHandler)
	mux.HandleFunc("/users", handler.CreateUserHandler)

	return mux
}
