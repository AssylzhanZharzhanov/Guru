package model

import "time"

type Mentor struct {
	ID        int       `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Info      string    `json:"info" db:"info"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
