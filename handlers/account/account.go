package account

import (
	"log"
	"net/http"

	"github.com/ibilalkayy/Bloop/website/handlers/callback"
	"github.com/ibilalkayy/Bloop/website/templates"
)

// func AccountPage(w http.ResponseWriter, r *http.Request) error {
// 	session, err := callback.Store.Get(r, "session-name")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return err
// 	}

// 	user, ok := session.Values["user"].(callback.User)
// 	if !ok {
// 		http.Redirect(w, r, "/login", http.StatusSeeOther)
// 		return nil
// 	}

// 	return templates.AccountTmpl.Execute(w, user)
// }

func AccountPage(w http.ResponseWriter, r *http.Request) error {
	log.Println("AccountPage called") // New logging line
	session, err := callback.Store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	user, ok := session.Values["user"].(callback.User)
	if !ok {
		log.Println("No user in session") // New logging line
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return nil
	}

	log.Printf("User in session: %+v", user) // New logging line
	return templates.AccountTmpl.Execute(w, user)
}
