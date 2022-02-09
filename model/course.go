package model

type Course struct {
	ID           int    `json:"id" db:"id"`
	MentorID     int    `json:"mentor_id,omitempty" db:"mentor_id"`
	CategoryID   int    `json:"category_id,omitempty" db:"category_id"`
	Name         string `json:"name,omitempty" db:"name"`
	Title        string `json:"title,omitempty" db:"title"`
	Description  string `json:"description,omitempty" db:"description"`
	Effect       string `json:"effect,omitempty" db:"effect"`
	IncludedData string `json:"included_data,omitempty" db:"included_data"`
	Price        int    `json:"price,omitempty" db:"price"`
	Views        int    `json:"views,omitempty" db:"views"`
	Purchases    int    `json:"purchases,omitempty" db:"purchases"`
	CreatedAt    int    `json:"created_at,omitempty" db:"created_at"`
}