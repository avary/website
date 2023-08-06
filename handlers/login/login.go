package login

import (
	"fmt"
	"net/http"

	"github.com/ibilalkayy/Bloop/website/auth"
)

func LoginPage(w http.ResponseWriter, r *http.Request) error {
	// Building the URL to Auth0's Universal Login
	redirectURL := fmt.Sprintf(
		"https://%s/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=openid profile email&prompt=login",
		auth.Auth0Domain, auth.Auth0ClientID, auth.Auth0CallbackURL)

	// Redirecting the user to Auth0's Universal Login
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)

	return nil
}
