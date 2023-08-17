package postgres

import "fmt"

func InsertIntoAccount(values []string) error {
	// Ensure the table exists
	db, err := ExecuteQueryFromSQLFile(1)
	if err != nil {
		return err
	}
	defer db.Close()

	// Validate that the values slice contains 3 strings
	if len(values) != 5 {
		return fmt.Errorf("Expected 5 values (names, usernames, emails, countries, shipping_addresses), got %d", len(values))
	}

	// Prepare the INSERT statement
	query := `
        INSERT INTO Account (names, usernames, emails, countries, shipping_addresses)
        VALUES ($1, $2, $3, $4, $5)
    `

	_, err = db.Exec(query, values[0], values[1], values[2], values[3], values[4])
	if err != nil {
		return fmt.Errorf("Error inserting value into Account: %v", err)
	}

	fmt.Println("Values inserted successfully!")
	return nil
}
