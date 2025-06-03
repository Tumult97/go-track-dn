package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

const connString = "server=localhost;database=tracker;trusted_connection=yes;encrypt=disable"

var DB *sql.DB

func Connect() {
	dsn := connString
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	var err error
	DB, err = sql.Open("sqlserver", dsn)
	if err != nil {
		log.Fatalf("Failed to create connection: %v", err)
	}

	err = DB.PingContext(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	fmt.Println("✅ Connected to SQL Server")
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
