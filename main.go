package main

import (
	"log"
	"net/http"

	"api/internal/database"
	"api/internal/httpapi"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := database.InitSchema(db); err != nil {
		log.Fatal(err)
	}

	router := httpapi.NewRouter(db)

	log.Fatal(http.ListenAndServe(":8000", router))
}
