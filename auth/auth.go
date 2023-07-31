package auth

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

var (
	Auth0Domain       string
	Auth0ClientID     string
	Auth0ClientSecret string
	CallbackURL       string
	Conf              *oauth2.Config
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Auth0Domain = os.Getenv("AUTH0_DOMAIN")
	Auth0ClientID = os.Getenv("AUTH0_CLIENT_ID")
	Auth0ClientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
	CallbackURL = os.Getenv("CALLBACK_URL")

	Conf = &oauth2.Config{
		ClientID:     Auth0ClientID,
		ClientSecret: Auth0ClientSecret,
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  fmt.Sprintf("https://%s/authorize", Auth0Domain),
			TokenURL: fmt.Sprintf("https://%s/oauth/token", Auth0Domain),
		},
		RedirectURL: CallbackURL,
	}
}
