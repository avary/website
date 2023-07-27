package authors

import (
	"net/http"

	"github.com/ibilalkayy/Bloop/website/templates"
)

func AuthorsSignupPage(w http.ResponseWriter, r *http.Request) error {
	return templates.AuthorsSignupTmpl.Execute(w, nil)
}
