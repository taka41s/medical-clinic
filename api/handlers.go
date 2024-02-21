package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func Root(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hello, world!",
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(responseData)
}

func Greet(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Greetings!",
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(responseData)
}
