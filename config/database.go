package config

import (
	"database/sql"
    "fmt"
    "log"
    "os"
	_ "github.com/lib/pq"
)

func SetupDatabase(){
    dbString := os.Getenv("GOOSE_DBSTRING")
    if dbString == "" {
        log.Fatal("GOOSE_DBSTRING environment variable is not set")
    }

    db, err := sql.Open("postgres", dbString)
    if err != nil {
        log.Fatalf("Failed to open database: %v", err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    fmt.Println("Successfully connected to the database")
}
