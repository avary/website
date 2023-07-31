package logout

import (
	"net/http"

	"github.com/ibilalkayy/Bloop/website/handlers/callback"
)

func LogoutPage(w http.ResponseWriter, r *http.Request) error {
	session, err := callback.Store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	// Delete the user value from the session
	delete(session.Values, "user")

	// Save the updated session
	session.Save(r, w)

	// Redirect to Auth0's logout endpoint
	http.Redirect(w, r, "http://localhost:8080", http.StatusTemporaryRedirect)
	return nil
}
