package main

import (
	"banking-api/internal/handler"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", handler.HelloHandler)
	http.HandleFunc("/health", handler.HealthHandler)
	http.HandleFunc("/users", handler.CreateUserHandler)

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		fmt.Println(err)
	}

}
