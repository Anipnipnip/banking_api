package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func PanggilHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	fmt.Fprintln(w, "Haloooo....", name)
}

func CekHandler(w http.ResponseWriter, r *http.Request) {

	cekRespon := map[string]string{
		"status":  "OK",
		"message": "Server berjalan",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(cekRespon)
}

func main() {

	http.HandleFunc("/hello", PanggilHandler)
	http.HandleFunc("/health", CekHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}

}
