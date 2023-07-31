package signup

import (
	"net/http"

	"github.com/ibilalkayy/Bloop/website/auth"
	"golang.org/x/oauth2"
)

func SignupPage(w http.ResponseWriter, r *http.Request) error {
	url := auth.Conf.AuthCodeURL("state", oauth2.SetAuthURLParam("screen_hint", "signup"), oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return nil
}
