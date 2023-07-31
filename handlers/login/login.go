package login

import (
	"net/http"

	"github.com/ibilalkayy/Bloop/website/auth"
	"golang.org/x/oauth2"
)

func LoginPage(w http.ResponseWriter, r *http.Request) error {
	url := auth.Conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	return nil
}
