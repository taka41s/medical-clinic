package api

import (
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func Root(w http.ResponseWriter, r *http.Request) {
	writeJSONResponse(w, "Hello, world!", http.StatusOK)
}

func Greet(w http.ResponseWriter, r *http.Request) {
	writeJSONResponse(w, "Greetings!", http.StatusOK)
}
