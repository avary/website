package books

import (
	"net/http"

	"github.com/ibilalkayy/Bloop/website/templates"
)

func BooksPage(w http.ResponseWriter, r *http.Request) error {
	return templates.BooksTmpl.Execute(w, nil)
}
