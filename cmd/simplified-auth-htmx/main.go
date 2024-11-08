package main

import (
	"log/slog"
	"os"

	"github.com/fearhunt/simplified-auth-htmx/internal/app/server"
)

func main() {
	session, err := server.SetupSessions()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	// Setup and start HTTP server
	server.StartServer(session)
}
