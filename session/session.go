package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

func SetSession(token, username, email string, w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user-session")
	session.Values["token"] = token
	session.Values["username"] = username
	session.Values["email"] = email
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

func GetUserData(r *http.Request) (string, string) {
	session, _ := store.Get(r, "user-session")
	username, _ := session.Values["username"].(string)
	email, _ := session.Values["email"].(string)
	return username, email
}