.PHONY: run
run:
	templ generate
	go run cmd/simplified-auth-htmx/main.go

build:
	go run buld cmd/simplified-auth-htmx/main.go

fmt:
	templ fmt ./internal/app/views
	go fmt ./...