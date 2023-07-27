package loops

import (
	"net/http"

	"github.com/ibilalkayy/Bloop/website/templates"
)

func LoopsPage(w http.ResponseWriter, r *http.Request) error {
	return templates.LoopsTmpl.Execute(w, nil)
}
