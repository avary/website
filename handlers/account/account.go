package account

import (
	"net/http"

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

	return templates.AccountTmpl.Execute(w, data)
}
