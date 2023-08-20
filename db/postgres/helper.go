package postgres

import "database/sql"

func DataExists(db *sql.DB, email, username string) (bool, error) {
	const query = "SELECT 1 FROM Account WHERE emails = $1 AND usernames = $2 LIMIT 1"

	var exists int
	err := db.QueryRow(query, email, username).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return exists == 1, nil
}
