package models

import (
	"context"
	"testing"
	"ajp-medical-clinic/config"
	"ajp-medical-clinic/models" // Import the package containing your User struct
    "github.com/volatiletech/sqlboiler/v4/boil"
)

func TestUserModel(t *testing.T) {
	// Initialize the model with sample data
	user := &models.User{
		Name:        "John Doe",
		Gender:      "male",
		Username:    "johndoe",
		Email:       "johndoe@example.com",
	}

	t.Run("TestGetters", func(t *testing.T) {
		if user.Name != "John Doe" {
			t.Errorf("Expected Name %s, got %s", "John Doe", user.Name)
		}
		if user.Username != "johndoe" {
			t.Errorf("Expected Username %s, got %s", "johndoe", user.Username)
		}
		if user.Email != "johndoe@example.com" {
			t.Errorf("Expected Email %s, got %s", "johndoe@example.com", user.Email)
		}
	})

	t.Run("TestSaveAndDelete", func(t *testing.T) {
		// Save the user to the database

		db, err := config.ConnectToDatabase()
		if err != nil {
			return
		}
		defer db.Close()

		err = user.Insert(context.Background(), db, boil.Infer())
		if err != nil {
			t.Fatalf("Error saving user to the database: %v", err)
		}
	})
}
