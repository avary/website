package sessions

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func SetSession(token string, w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user-session")
	session.Values["token"] = token
	session.Save(r, w)
}

func GetSession(r *http.Request) string {
	session, _ := store.Get(r, "user-session")
	if token, ok := session.Values["token"].(string); ok {
		return token
	}
	return ""
}

func ClearSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user-session")
	session.Options.MaxAge = -1 // deletes the session
	session.Save(r, w)
}
