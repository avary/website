package account

import (
	"fmt"
	"net/http"

	"github.com/ibilalkayy/Bloop/website/db/postgres"
	"github.com/ibilalkayy/Bloop/website/model"
	"github.com/ibilalkayy/Bloop/website/session"
	"github.com/ibilalkayy/Bloop/website/templates"
)

type AccountData struct {
	Username string
	Email    string
}

func AccountPage(w http.ResponseWriter, r *http.Request) error {
	username, email := session.GetUserData(r)
	data := AccountData{
		Username: username,
		Email:    email,
	}

	// Ensure the table exists
	db, err := postgres.ExecuteQueryFromSQLFile(1) // Assuming the create table query is the first in the file
	if err != nil {
		return fmt.Errorf("Failed to ensure table: %v", err)
	}
	defer db.Close()

	if r.Method == "GET" {
		return templates.AccountTmpl.Execute(w, data)
	} else if r.Method == "POST" {
		accountDetails := model.AccountDetails{
			Name:            r.FormValue("name"),
			Username:        username,
			Email:           email,
			Country:         r.FormValue("country"),
			ShippingAddress: r.FormValue("shipping-address"),
		}
		values := map[string]string{
			"names":              accountDetails.Name,
			"usernames":          accountDetails.Username,
			"emails":             accountDetails.Email,
			"countries":          accountDetails.Country,
			"shipping_addresses": accountDetails.ShippingAddress,
		}

		exists, err := postgres.DataExists(db, email, username)
		if err != nil {
			return fmt.Errorf("Failed to check if data exists: %v", err)
		}

		if exists {
			err = postgres.UpdateAccount(email, username, values)
			if err != nil {
				return fmt.Errorf("Failed to update Account: %v", err)
			}
		} else {
			// Converting map values to a slice for InsertIntoAccount function
			valuesSlice := []string{values["names"], values["usernames"], values["emails"], values["countries"], values["shipping_addresses"]}
			err = postgres.InsertIntoAccount(valuesSlice)
			if err != nil {
				return fmt.Errorf("Failed to insert values: %v", err)
			}
		}

		return templates.AccountTmpl.Execute(w, data)
	}

	return nil
}
