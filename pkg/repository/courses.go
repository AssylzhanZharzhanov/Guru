package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/zharzhanov/myguru/model"
)

type CoursesRepository struct {
	db *sqlx.DB
}

func (r *CoursesRepository) GetCourses() ([]model.Course, error) {
	courses := make([]model.Course, 0)

	query := fmt.Sprintf(`
		SELECT * FROM %s
	`, coursesTable)

	err := r.db.Select(&courses, query)
	if err != nil {
		return courses, err
	}

	return courses, err
}

func (r *CoursesRepository) GetTrendingCourses() ([]model.Course, error) {
	panic("implement me")
}

func (r *CoursesRepository) GetComingSoonCourses() ([]model.Course, error) {
	panic("implement me")
}

func (r *CoursesRepository) GetCourseByID(id int) (model.Course, error) {
	panic("implement me")
}

func NewCoursesRepository(db *sqlx.DB) *CoursesRepository {
	return &CoursesRepository{db: db}
}
