package postgres

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectToDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbHost, dbPort)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Cannot open database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Cannot ping the database: %v", err)
	}

	return db, nil
}

func ExecuteSQLQuery(queryNumber int) error {
	// Connect to the database
	db, err := ConnectToDB()
	if err != nil {
		return fmt.Errorf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Read the content of the SQL file
	content, err := ioutil.ReadFile("db/postgres/db.SQL")
	if err != nil {
		return fmt.Errorf("Error reading SQL file: %v", err)
	}

	// Split the content by semicolon to get individual SQL statements
	queries := strings.Split(string(content), ";")

	if queryNumber <= 0 || queryNumber > len(queries) {
		return fmt.Errorf("Invalid query number. There are %d queries in the file", len(queries))
	}

	query := strings.TrimSpace(queries[queryNumber-1])
	if query == "" {
		return fmt.Errorf("Query %d is empty", queryNumber)
	}

	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("Error executing query %d: %v", queryNumber, err)
	}
	fmt.Printf("Executed query %d successfully!\n", queryNumber)

	return nil
}
