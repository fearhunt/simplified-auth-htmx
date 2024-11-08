package server

import (
	"net/http"
	"time"

	"github.com/golangcollege/sessions"
	gsessions "github.com/gorilla/sessions"
)

func SetupSessions() (*sessions.Session, error) {
	var session *sessions.Session
	var secret = []byte("super-secure-secret")
	session = sessions.New(secret)
	session.Lifetime = 1 * time.Hour

	key := "Secret-session-key"
	maxAge := 3600 * 1 // 1 hour

	cookieStore := gsessions.NewCookieStore([]byte(key))
	cookieStore.MaxAge(maxAge)

	cookieStore.Options.Domain = "localhost"
	cookieStore.Options.Path = ""
	cookieStore.Options.HttpOnly = true // HttpOnly should always be enabled
	cookieStore.Options.Secure = true

	return session, nil
}

func logoutHandler(session *sessions.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session.Destroy(r)
		w.Header().Set("HX-Redirect", "/") // Use the special HTMX redirect header to trigger a full page reload.
	}
}
