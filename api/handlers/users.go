package handlers

import (
	"context"
	"net/http"
	"ajp-medical-clinic/models"
	"ajp-medical-clinic/config"
	"encoding/json"
	"github.com/volatiletech/sqlboiler/v4/boil"
    "github.com/volatiletech/sqlboiler/v4/queries/qm"
    "github.com/gorilla/mux"

)

func FetchUsers(w http.ResponseWriter, r *http.Request) {
    // Connect to the database
    db, err := config.ConnectToDatabase()
    if err != nil {
        sendJSONResponse(w, nil, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Fetch all users from the database
    users, err := models.Users().All(context.Background(), db)
    if err != nil {
        sendJSONResponse(w, nil, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Send JSON response with data and optional message
    sendJSONResponse(w, users, "Users fetched successfully", http.StatusOK)
}


func RegisterUser(w http.ResponseWriter, r *http.Request) {
    // Parse request body into a User struct
    var newUser models.User
    err := json.NewDecoder(r.Body).Decode(&newUser)
    if err != nil {
        sendJSONResponse(w, nil, "Bad Request", http.StatusBadRequest)
        return
    }

    // Connect to the database
    db, err := config.ConnectToDatabase()
    if err != nil {
        sendJSONResponse(w, nil, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Create a new User model instance
    user := models.User{
        Name:        newUser.Name,
        Username:    newUser.Username,
        Email:       newUser.Email,
        Gender:      newUser.Gender,
    }

    err = user.Insert(context.Background(), db, boil.Infer())
    if err != nil {
        sendJSONResponse(w, nil, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    sendJSONResponse(w, nil, "User created successfully", http.StatusCreated)
}

func FetchUserByID(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the URL path parameter
	vars := mux.Vars(r)
	userID := vars["id"]
	if userID == "" {
		sendJSONResponse(w, nil, "User ID is required", http.StatusBadRequest)
		return
	}

	// Connect to the database
	db, err := config.ConnectToDatabase()
	if err != nil {
		sendJSONResponse(w, nil, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Fetch the user from the database by ID
	user, err := models.Users(qm.Where("id = ?", userID)).One(context.Background(), db)
	if err != nil {
		sendJSONResponse(w, nil, "User not found", http.StatusNotFound)
		return
	}

	// Send JSON response with the fetched user
	sendJSONResponse(w, user, "User fetched successfully", http.StatusOK)
}




