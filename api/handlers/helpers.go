package handlers

import (
	"encoding/json"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, message string, statusCode int) {
	response := Response{
		Message: message,
	}
	responseData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseData)
}
