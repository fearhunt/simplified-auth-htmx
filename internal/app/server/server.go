package server

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/fearhunt/simplified-auth-htmx/internal/app/views"
	"github.com/golangcollege/sessions"

	// "github.com/fearhunt/simplified-auth-htmx/internal/app/store"
	// "github.com/fearhunt/simplified-auth-htmx/internal/app/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartServer(session *sessions.Session) {
	// Set-up chi router with middleware
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(session.Enable)

	// Page specific handlers
	r.Get("/", indexPage(session))
	// r.Get("/", templ.Handler(views.Index(session)).ServeHTTP)

	// Start plain HTTP listener
	fmt.Println("starting server simplified-auth-htmx...")
	_ = http.ListenAndServe(":3000", r)
}

func indexPage(session *sessions.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var a string = session.GetString(r, "astaga")
		fmt.Printf("%v", a)

		templ.Handler(views.Index(a)).ServeHTTP(w, r)

		return
	}
}
