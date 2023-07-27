package signup

import (
	"net/http"

	"github.com/ibilalkayy/Bloop/website/templates"
)

func SignupPage(w http.ResponseWriter, r *http.Request) error {
	return templates.SignupTmpl.Execute(w, nil)
}
