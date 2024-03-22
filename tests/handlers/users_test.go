package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"ajp-medical-clinic/api/handlers"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	// Create a request body with user data
	userData := map[string]interface{}{
		"name":     "John Doe",
		"username": "johndoe",
		"email":    "johndoe@example.com",
		"gender":   "male",
	}
	reqBody, err := json.Marshal(userData)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// Create a mock HTTP request
	
	req := httptest.NewRequest("POST", "/user", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a mock HTTP response recorder
	w := httptest.NewRecorder()

	// Call the handler function with the mock request and response recorder
	handlers.RegisterUser(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// Parse the response body into a map
	var responseBody map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	// Check the response message
	assert.Equal(t, "User created successfully", responseBody["message"])
}
