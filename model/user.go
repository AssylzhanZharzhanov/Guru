package model

import "time"

type User struct {
	ID         int       `json:"id" db:"id"`
	Email      string    `json:"email" db:"email"`
	Password   string    `json:"password" db:"password"`
	FirstName  string    `json:"first_name" db:"first_name"`
	LastName   string    `json:"last_name" db:"last_name"`
	Title      string    `json:"title" db:"title"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}