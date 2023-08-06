package logout

import (
	"net/http"

	"github.com/ibilalkayy/Bloop/website/templates"
)

func LogoutPage(w http.ResponseWriter, r *http.Request) error {
	return templates.LogoutTmpl.Execute(w, nil)
}
