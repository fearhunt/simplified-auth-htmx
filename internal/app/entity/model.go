package entity

import "time"

type User struct {
	ID        string
	Username  string
	Password  string // you read it right!
	Name      string
	Email     string
	CreatedAt time.Time
}
