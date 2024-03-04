package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "medical_clinic"
)

func SetupDatabase() error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable",
		host, port, user, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", dbname).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		_, err := db.Exec("CREATE DATABASE " + dbname)
		if err != nil {
			return err
		}
		fmt.Printf("Database '%s' created successfully.\n", dbname)
	} else {
		fmt.Printf("Database '%s' already exists.\n", dbname)
	}

	err = goose.Up(db, "./config/migrations")
	if err != nil {
		return err
	}

	return nil
}
