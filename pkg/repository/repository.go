package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {

}

type Courses interface {
	GetCourses()
}

type Repository struct {
	Authorization
	Courses
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Courses: NewCoursesRepository(db),
	}
}