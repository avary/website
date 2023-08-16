package postgres

import "fmt"

func InsertIntoAccount(values []string) error {
	// Ensure the table exists
	db, err := ExecuteQueryFromSQLFile(1) // Assuming the create table query is the first in the file
	if err != nil {
		return err
	}
	defer db.Close()

	// Validate that the values slice contains 3 strings
	if len(values) != 3 {
		return fmt.Errorf("Expected 3 values (names, countries, shipping_addresses), got %d", len(values))
	}

	// Prepare the INSERT statement
	query := `
        INSERT INTO Account (names, countries, shipping_addresses)
        VALUES ($1, $2, $3)
    `

	_, err = db.Exec(query, values[0], values[1], values[2])
	if err != nil {
		return fmt.Errorf("Error inserting value into Account: %v", err)
	}

	fmt.Println("Values inserted successfully!")
	return nil
}
