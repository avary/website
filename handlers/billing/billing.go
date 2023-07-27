package billing

import (
	"net/http"

	"github.com/ibilalkayy/Bloop/website/templates"
)

func BillingPage(w http.ResponseWriter, r *http.Request) error {
	return templates.BillingTmpl.Execute(w, nil)
}
