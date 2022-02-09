package repository

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/zharzhanov/myguru/model"
)

type CoursesRepository struct {
	db *sqlx.DB
}

func (r *CoursesRepository) GetCourses() ([]model.Course, error) {
	panic("implement me")
}

func NewCoursesRepository(db *sqlx.DB) *CoursesRepository {
	return &CoursesRepository{db: db}
}
