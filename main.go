package main

import (
	"ajp-medical-clinic/api"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	defer fmt.Println("End of program")
	defer os.Exit(1)

	startServer()
}

func startServer() {
	http.HandleFunc("/", api.Root)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
