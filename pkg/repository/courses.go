package repository

import "github.com/jmoiron/sqlx"

type CoursesRepository struct {
	db *sqlx.DB
}

func (r *CoursesRepository) GetCourses() {
	panic("implement me")
}

func NewCoursesRepository(db *sqlx.DB) *CoursesRepository {
	return &CoursesRepository{db: db}
}
