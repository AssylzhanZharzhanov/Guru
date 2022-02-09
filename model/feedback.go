package model

import "time"

type Feedback struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	CourseID  int       `json:"course_id" db:"course_id"`
	Comment   string    `json:"comment" db:"comment"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
