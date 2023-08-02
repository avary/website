package callback

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/ibilalkayy/Bloop/website/auth"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Pic   string `json:"picture"`
}

var Store = sessions.NewCookieStore([]byte("1234"))

func CallbackPage(w http.ResponseWriter, r *http.Request) {
	log.Println("CallbackPage called") // New logging line
	code := r.URL.Query().Get("code")

	token, err := auth.Conf.Exchange(context.Background(), code)
	if err != nil {
		log.Println("Error exchanging code:", err) // New logging line
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := http.Get("https://" + auth.Auth0Domain + "/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Println("Error getting user info:", err) // New logging line
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)
	var user User
	json.Unmarshal(data, &user)

	session, err := Store.Get(r, "session-name")
	if err != nil {
		log.Println("Error getting session:", err) // New logging line
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["user"] = user
	if err := session.Save(r, w); err != nil { // Here we handle any errors from saving the session
		log.Println("Error saving session:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Redirecting to home") // New logging line
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
