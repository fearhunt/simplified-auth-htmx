package entity

import "time"

type User struct {
	ID        string
	Username  string
	Password  string
	Name      string
	Email     string
	CreatedAt time.Time
}

type FailedLoginAttempts struct {
	UsernameCheck int
	Password      int
}

type ValidatePasswordInput struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	CorrectPassword string `json:"correct_password"`
	IsMasked        bool   `json:"masked"`
}
