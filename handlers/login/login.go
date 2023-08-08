package login

import (
	"fmt"
	"net/http"

	"github.com/ibilalkayy/Bloop/website/auth"
	"github.com/ibilalkayy/Bloop/website/session"
)

func LoginPage(w http.ResponseWriter, r *http.Request) error {
	// Check if user is already logged in
	token := session.GetSession(r)
	if token != "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return nil
	}

	// If not, proceed to Auth0 login
	redirectURL := fmt.Sprintf(
		"https://%s/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=openid profile email&prompt=login",
		auth.Auth0Domain, auth.Auth0ClientID, auth.Auth0CallbackURL)

	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
	return nil
}
