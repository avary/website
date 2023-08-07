package auth

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Auth0Domain       string
	Auth0ClientID     string
	Auth0ClientSecret string
	Auth0CallbackURL  string
)

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Auth0Domain = os.Getenv("AUTH0_DOMAIN")
	Auth0ClientID = os.Getenv("AUTH0_CLIENT_ID")
	Auth0ClientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
	Auth0CallbackURL = os.Getenv("CALLBACK_URL")
}
