package handler

import (
	"banking-api/internal/model"
	"banking-api/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	fmt.Fprintln(w, "Haloooo..", name)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	response := map[string]string{
		"status":  "OK",
		"message": "Banking is running",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User

	json.NewDecoder(r.Body).Decode(&user)

	user = service.CreateUser(user)

	fmt.Fprintln(w, "Masukkan nama", user.Name)
	fmt.Fprintln(w, "Masukkan email", user.Email)
}
