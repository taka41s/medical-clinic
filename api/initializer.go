package api

import (
	"net/http"
	"log"
)

func StartServer() {
	for _, route := range routes {
		http.HandleFunc(route.Path, route.Handler)
	}

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
