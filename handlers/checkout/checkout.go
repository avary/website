package checkout

import (
	"net/http"

	"github.com/ibilalkayy/Bloop/website/templates"
)

func CheckoutPage(w http.ResponseWriter, r *http.Request) error {
	return templates.CheckoutTmpl.Execute(w, nil)
}
