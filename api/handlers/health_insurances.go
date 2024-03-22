package handlers

import (
	"context"
	"net/http"
	"ajp-medical-clinic/models"
	"ajp-medical-clinic/config"
)

func FetchHealthInsurances(w http.ResponseWriter, r *http.Request) {
    // Connect to the database
    db, err := config.ConnectToDatabase()
    if err != nil {
        sendJSONResponse(w, nil, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Fetch all e from the database
    e, err := models.HealthInsurances().All(context.Background(), db)
    if err != nil {
        sendJSONResponse(w, nil, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Send JSON response with data and optional message
    sendJSONResponse(w, e, "Health insurances fetched successfully", http.StatusOK)
}
