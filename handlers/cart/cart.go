package cart

import (
	"net/http"

	"github.com/ibilalkayy/Bloop/website/templates"
)

func CartPage(w http.ResponseWriter, r *http.Request) error {
	return templates.CartTmpl.Execute(w, nil)
}
