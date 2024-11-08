.PHONY: run
run:
	templ generate
	go run cmd/simplified-auth-htmx/main.go

run-dev:
	templ generate --watch --cmd="go run cmd/simplified-auth-htmx/main.go"

build:
	go build cmd/simplified-auth-htmx/main.go

fmt:
	templ fmt ./internal/app/views
	go fmt ./...