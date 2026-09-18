package main

import (
	"banking-api/internal/database"
	"banking-api/routes"
	"fmt"
	"net/http"
)

func main() {

	db, err := database.ConnectDB()

	if err != nil {
		fmt.Println("Gagal konek ke database : ", err)
		return
	}

	fmt.Println("Database berhasil terhubung")

	_ = db

	router := routes.SetupRoutes()

	err = http.ListenAndServe(":8000", router)

	if err != nil {
		fmt.Println(err)
	}

}
