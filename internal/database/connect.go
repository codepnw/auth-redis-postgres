package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDatabase(dbURL string) *sql.DB {
	if dbURL == "" {
		log.Fatal("DB_URL not found or empty")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	var testQuery int
	if err := conn.QueryRow("SELECT 1").Scan(&testQuery); err != nil {
		log.Fatal("database connection test failed")
	}
	log.Println("database connection successfully")

	return conn
}
