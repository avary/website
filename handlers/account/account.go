package account

import (
	"log"
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
		values := []string{accountDetails.Name, accountDetails.Username, accountDetails.Email, accountDetails.Country, accountDetails.ShippingAddress}
		err := postgres.InsertIntoAccount(values)
		if err != nil {
			log.Fatalf("Failed to insert values: %v", err)
		}
		return templates.AccountTmpl.Execute(w, data)
	}
	return nil
}
