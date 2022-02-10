package repository

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/zharzhanov/myguru/model"
)

type CoursesRepository struct {
	db *sqlx.DB
}

func (r *CoursesRepository) GetTrendingCourses() ([]model.Course, error) {
	panic("implement me")
}

func (r *CoursesRepository) GetComingSoonCourses() ([]model.Course, error) {
	panic("implement me")
}

func (r *CoursesRepository) GetCourses() ([]model.Course, error) {
	panic("implement me")
}

func (r *CoursesRepository) GetCourseByID(id int) (model.Course, error) {
	panic("implement me")
}

func NewCoursesRepository(db *sqlx.DB) *CoursesRepository {
	return &CoursesRepository{db: db}
}
