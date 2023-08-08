package callback

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/ibilalkayy/Bloop/website/auth"
	"github.com/ibilalkayy/Bloop/website/session"
)

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	IDToken     string `json:"id_token"`
	Scope       string `json:"scope"`
	ExpiresIn   int    `json:"expires_in"`
}

func CallbackPage(w http.ResponseWriter, r *http.Request) error {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "No code returned from Auth0", http.StatusBadRequest)
		return nil
	}

	tokenResp, err := exchangeCodeForToken(code)
	if err != nil {
		http.Error(w, "Error exchanging code for token", http.StatusInternalServerError)
		return err
	}

	// Store the token in the session
	session.SetSession(tokenResp.AccessToken, w, r)

	// Here, you'd store the token securely, then redirect to the desired page
	// For simplicity, we're just redirecting
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}

func exchangeCodeForToken(code string) (*tokenResponse, error) {
	// Construct the request payload
	payload := url.Values{}
	payload.Set("grant_type", "authorization_code")
	payload.Set("client_id", auth.Auth0ClientID)
	payload.Set("client_secret", auth.Auth0ClientSecret)
	payload.Set("code", code)
	payload.Set("redirect_uri", "http://localhost:8080/callback")

	resp, err := http.PostForm("https://"+auth.Auth0Domain+"/oauth/token", payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var token tokenResponse
	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
