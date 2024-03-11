package api

import (
	"net/http"
	"log"
)

func StartServer() {
	InitializeRoutes()

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
