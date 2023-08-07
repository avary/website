package middleware

import (
	"log"
	"net/http"

	"github.com/ibilalkayy/Bloop/website/sessions"
)

// Type to be used as a parameter in a function
type MyHandlerFunc func(w http.ResponseWriter, r *http.Request) error

// ErrorHandling handles the error and returns http.Handler
func ErrorHandling(h MyHandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})
}

func AuthenticatedOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := sessions.GetSession(r)
		if token == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	}
}
