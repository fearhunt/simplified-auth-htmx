package server

import (
	"encoding/gob"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/fearhunt/simplified-auth-htmx/internal/app/entity"
	"github.com/fearhunt/simplified-auth-htmx/internal/app/views"
	"github.com/golangcollege/sessions"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartServer(session *sessions.Session) {
	// Register the type with the encoding/gob package
	gob.Register(entity.FailedLoginAttempts{})

	// Set-up chi router with middleware
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(session.Enable)

	// Page specific handlers
	r.Get("/", indexPage(session))

	// API handlers
	r.Get("/user/find", findUsername(session))

	// r.Route("/api", func(r chi.Router) {
	// 	r.Get("/user/find", findUsername())
	// })

	// Start plain HTTP listener
	fmt.Println("starting server simplified-auth-htmx...")
	_ = http.ListenAndServe(":8215", r)
}

func indexPage(session *sessions.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fla := GetFailedLoginAttempts(session, r)
		templ.Handler(views.Index(nil, fla)).ServeHTTP(w, r)
	}
}
