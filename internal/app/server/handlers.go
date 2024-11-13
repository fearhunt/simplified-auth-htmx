package server

import (
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/fearhunt/simplified-auth-htmx/internal/app/entity"
	"github.com/fearhunt/simplified-auth-htmx/internal/app/views"
	"github.com/golangcollege/sessions"
)

func getAllUser() []entity.User {
	// yes, hardcode
	users := []entity.User{
		{
			ID:        "1",
			Username:  "john_doe",
			Password:  "StrongPassword123",
			Name:      "John Doe",
			Email:     "john.doe@example.com",
			CreatedAt: time.Date(2023, time.July, 14, 8, 30, 0, 0, time.UTC),
		},
		{
			ID:        "2",
			Username:  "jane_smith",
			Password:  "SecurePassword2024",
			Name:      "Jane Smith",
			Email:     "jane.smith@example.com",
			CreatedAt: time.Date(2022, time.August, 22, 12, 0, 0, 0, time.UTC),
		},
		{
			ID:        "3",
			Username:  "admin_user",
			Password:  "AdminPassword2024",
			Name:      "Admin User",
			Email:     "admin@example.com",
			CreatedAt: time.Date(2021, time.March, 5, 14, 45, 0, 0, time.UTC),
		},
		{
			ID:        "4",
			Username:  "michael_brown",
			Password:  "MichaelBrown2023",
			Name:      "Michael Brown",
			Email:     "michael.brown@example.com",
			CreatedAt: time.Date(2020, time.September, 11, 9, 15, 0, 0, time.UTC),
		},
		{
			ID:        "5",
			Username:  "lucy_white",
			Password:  "LucyWhitePassword24",
			Name:      "Lucy White",
			Email:     "lucy.white@example.com",
			CreatedAt: time.Date(2024, time.June, 18, 17, 30, 0, 0, time.UTC),
		},
	}

	return users
}

func findUsername(session *sessions.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		qUser := r.URL.Query().Get("username")
		users := getAllUser()

		for i := range users {
			if users[i].Username == qUser {
				currentUser := users[i]
				templ.Handler(views.Login(&currentUser)).ServeHTTP(w, r)
				// session.Put(r, "name", users[i].Name)
				// session.Put(r, "username", users[i].Username)
				// session.Put(r, "password", users[i].Password)

				// templ.Handler(views.Login(users[i].Username, users[i].Password, users[i].Name)).ServeHTTP(w, r)
				return
			}
		}

		session.Destroy(r)
		templ.Handler(views.Login(&entity.User{
			Username: qUser,
		})).ServeHTTP(w, r)
	}
}
