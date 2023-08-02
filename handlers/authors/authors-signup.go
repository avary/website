package authors

import (
	"net/http"

	"github.com/ibilalkayy/Bloop/website/handlers/callback"
	"github.com/ibilalkayy/Bloop/website/templates"
)

func AuthorsSignupPage(w http.ResponseWriter, r *http.Request) error {
	session, err := callback.Store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	user, ok := session.Values["user"].(callback.User)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return nil
	}
	return templates.AuthorsSignupTmpl.Execute(w, user)
}
