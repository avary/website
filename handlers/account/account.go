package account

import (
	"net/http"

	"github.com/ibilalkayy/Bloop/website/templates"
)

func AccountPage(w http.ResponseWriter, r *http.Request) error {
	return templates.AccountTmpl.Execute(w, nil)
}
