package signup

import (
	"fmt"
	"net/http"

	"github.com/ibilalkayy/Bloop/website/auth"
)

func SignupPage(w http.ResponseWriter, r *http.Request) error {
	// Constructing the URL
	redirectURL := fmt.Sprintf(
		"https://%s/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=openid profile email&screen_hint=signup",
		auth.Auth0Domain, auth.Auth0ClientID, auth.Auth0CallbackURL)

	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
	return nil
}
