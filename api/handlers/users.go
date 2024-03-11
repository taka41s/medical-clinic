package handlers

import (
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	WriteJSONResponse(w, "Hello, world!", http.StatusOK)
}

func FetchUsers(w http.ResponseWriter, r *http.Request) {
	WriteJSONResponse(w, "Greetings!", http.StatusOK)
}
