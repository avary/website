package home

import (
	"net/http"

	"github.com/ibilalkayy/Bloop/website/templates"
)

func HomePage(w http.ResponseWriter, r *http.Request) error {
	return templates.HomeTmpl.Execute(w, nil)
}
