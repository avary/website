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

	db, err := postgres.ConnectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	return templates.AccountTmpl.Execute(w, data)
}
