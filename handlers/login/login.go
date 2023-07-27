package login

import (
	"net/http"

	"github.com/ibilalkayy/Bloop/website/templates"
)

func LoginPage(w http.ResponseWriter, r *http.Request) error {
	return templates.LoginTmpl.Execute(w, nil)
}
