package account

import (
	"log"
	"net/http"

	"github.com/ibilalkayy/Bloop/website/db/postgres"
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

	err := postgres.ExecuteSQLQuery(1)
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}

	return templates.AccountTmpl.Execute(w, data)
}
