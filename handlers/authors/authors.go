package authors

import (
	"net/http"

	"github.com/ibilalkayy/Bloop/website/templates"
)

func AuthorsPage(w http.ResponseWriter, r *http.Request) error {
	return templates.AuthorsTmpl.Execute(w, nil)
}
