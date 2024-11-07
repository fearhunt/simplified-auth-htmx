package server

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/fearhunt/simplified-auth-htmx/internal/app/views"

	// "github.com/fearhunt/simplified-auth-htmx/internal/app/store"
	// "github.com/fearhunt/simplified-auth-htmx/internal/app/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func StartServer() {
	// Set-up chi router with middleware
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Page specific handlers
	r.Get("/", templ.Handler(views.Index()).ServeHTTP)

	// Start plain HTTP listener
	fmt.Println("starting server simplified-auth-htmx...")
	_ = http.ListenAndServe(":3000", r)
}
