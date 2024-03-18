package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "medical_clinic"
)

func SetupDatabase() error {
	db, err := ConnectToDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := checkIfDatabaseExists(db, dbname); err != nil {
		return err
	}

	return nil
}

func ConnectToDatabase() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}


func checkIfDatabaseExists(db *sql.DB, dbname string) error {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", dbname).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		if err := createDatabase(db, dbname); err != nil {
			return err
		}
	} else {
		fmt.Printf("Database '%s' already exists.\n", dbname)
	}

	return nil
}

func createDatabase(db *sql.DB, dbname string) error {
	_, err := db.Exec("CREATE DATABASE " + dbname)
	if err != nil {
		return err
	}
	fmt.Printf("Database '%s' created successfully.\n", dbname)
	return nil
}
