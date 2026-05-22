package auth

import (
	"log"
	"os"

	"github.com/gorilla/sessions"
	_ "github.com/joho/godotenv/autoload"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

var key string = os.Getenv("SESSION_KEY")

const (
	MaxAge = 86400 * 30
	IsProd = false
)

func NewAuth() {
	googleClientID, ok := os.LookupEnv("GOOGLE_OAUTH_CLIENT_ID")
	if !ok {
		log.Fatal("no google client id provided")
	}

	googleSecret, ok := os.LookupEnv("GOOGLE_OAUTH_SECRET")
	if !ok {
		log.Fatal("no google secret provided")
	}

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	// store.Options.SameSite = http.SameSiteStrictMode
	store.Options.Secure = IsProd

	gothic.Store = store

	goth.UseProviders(
		google.New(googleClientID, googleSecret, "http://localhost:8080/api/auth/google/callback"),
	)
}
