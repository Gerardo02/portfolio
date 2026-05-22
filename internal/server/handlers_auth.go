package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

func (s *Server) callbackHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Println(user.UserID)

	http.Redirect(w, r, "/", http.StatusFound)
}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
	// try to get the user without re-authenticating
	if _, err := gothic.CompleteUserAuth(w, r); err == nil {
		fmt.Println("ya estabas logueado compa")
		http.Redirect(w, r, "/", http.StatusFound)
		return
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
	gothic.Logout(w, r)

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
