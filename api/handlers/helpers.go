package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
    Message *string      `json:"message,omitempty"`
    Data    interface{} `json:"data,omitempty"`
}

func sendJSONResponse(w http.ResponseWriter, data interface{}, message interface{}, statusCode int) {
    response := Response{
        Data: data,
    }

    switch msg := message.(type) {
    case string:
        if msg != "" {
            response.Message = &msg
        }
    case *string:
        if msg != nil && *msg != "" {
            response.Message = msg
        }
    }

    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    w.Write(jsonResponse)
}

