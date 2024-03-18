package handlers

import (
	"context"
	"net/http"
	"ajp-medical-clinic/models"
	"ajp-medical-clinic/config"
	"encoding/json"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Response struct {
	Message string `json:"message"`
}

func FetchUsers(w http.ResponseWriter, r *http.Request) {
	// Connect to the database
	db, err := config.ConnectToDatabase()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Fetch all users from the database
	users, err := models.Users().All(context.Background(), db)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Convert users to JSON
	usersJSON, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set response headers and write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usersJSON)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Parse request body into a User struct
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		WriteJSONResponse(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Connect to the database
	db, err := config.ConnectToDatabase()
	if err != nil {
		WriteJSONResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Create a new User model instance
	user := models.User{
		Username:    newUser.Username,
		Email:       newUser.Email,
		Gender:      newUser.Gender,
	}

	err = user.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		WriteJSONResponse(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	WriteJSONResponse(w, "User created successfully", http.StatusCreated)
}
