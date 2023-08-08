package logout

import (
	"fmt"
	"net/http"

	"github.com/ibilalkayy/Bloop/website/auth"
	"github.com/ibilalkayy/Bloop/website/session"
)

func LogoutPage(w http.ResponseWriter, r *http.Request) error {
	// Clear the session
	session.ClearSession(w, r)

	// Clear the Auth0 session by redirecting to the Auth0 logout endpoint
	redirectURL := fmt.Sprintf("https://%s/v2/logout?client_id=%s&returnTo=%s",
		auth.Auth0Domain, auth.Auth0ClientID, "http://localhost:8080/")

	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
	return nil
}
