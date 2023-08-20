package postgres

import (
	"fmt"
	"strings"
)

func UpdateAccount(email, username string, newValues map[string]string) error {
	// Ensure the table exists
	db, err := ExecuteQueryFromSQLFile(1) // Assuming the create table query is the first in the file
	if err != nil {
		return err
	}
	defer db.Close()

	// Building the SET clause dynamically based on provided fields
	var setClauses []string
	var args []interface{}

	i := 1
	for key, value := range newValues {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", key, i))
		args = append(args, value)
		i++
	}
	args = append(args, email, username)

	query := fmt.Sprintf(
		"UPDATE Account SET %s WHERE emails = $%d AND usernames = $%d",
		strings.Join(setClauses, ", "),
		i,
		i+1,
	)

	_, err = db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("Error updating Account: %v", err)
	}

	fmt.Println("Account updated successfully!")
	return nil
}
