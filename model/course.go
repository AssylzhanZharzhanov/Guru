package model

import "time"

type Course struct {
	ID           int          `json:"id" db:"id"`
	StatusID     int          `json:"status_id" db:"status_id"`
	MentorID     int          `json:"mentor_id" db:"mentor_id"`
	CategoryID   int          `json:"category_id" db:"category_id"`
	Name         string       `json:"name" db:"name"`
	Title        string       `json:"title" db:"title"`
	Description  string       `json:"description" db:"description"`
	Effect       string       `json:"effect" db:"effect"`
	IncludedData string       `json:"included_data" db:"included_data"`
	Price        int          `json:"price" db:"price"`
	Views        int          `json:"views" db:"views"`
	Purchases    int          `json:"purchases" db:"purchases"`
	CreatedAt    time.Time    `json:"created_at" db:"created_at"`
}