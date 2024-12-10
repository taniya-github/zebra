package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Database connection details
	dbUser := "testuser"
	dbPassword := "testpassword"
	dbName := "testdb"
	dbHost := "db" // Docker service name from docker-compose

	// MySQL DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3307)/%s", dbUser, dbPassword, dbHost, dbName)

	// Connect to MySQL
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}

	// Set up HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Connection successful!"))
	})

	log.Println("Server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

