package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/zharzhanov/myguru/model"
)

type CoursesRepository struct {
	db *sqlx.DB
}

func (r *CoursesRepository) GetCourses(lang string) ([]model.Course, error) {
	courses := make([]model.Course, 0)

	query := fmt.Sprintf(`
		SELECT
			id,
			status_id,
			mentor_id,
			category_id,
			name,
			title,
			description,
			effect,
			included_data,
			price,
			views,
			purchases,
			created_at
		FROM %s
		ORDER BY id DESC
	`, coursesTable)

	err := r.db.Select(&courses, query)
	if err != nil {
		return courses, err
	}

	return courses, err
}

func (r *CoursesRepository) GetTrendingCourses(lang string) ([]model.Course, error) {
	courses := make([]model.Course, 0)

	query := fmt.Sprintf(`
		SELECT
			id,
			status_id,
			mentor_id,
			category_id,
			name,
			title,
			description,
			effect,
			included_data,
			price,
			views,
			purchases,
			created_at
		FROM %s
		ORDER BY purchases DESC, views DESC
	`, coursesTable)

	err := r.db.Select(&courses, query)
	if err != nil {
		return courses, err
	}

	return courses, err
}

func (r *CoursesRepository) GetComingSoonCourses(lang string) ([]model.Course, error) {
	courses := make([]model.Course, 0)

	query := fmt.Sprintf(`
		SELECT
			id,
			status_id,
			mentor_id,
			category_id,
			name,
			title,
			description,
			effect,
			included_data,
			price,
			views,
			purchases,
			created_at
		FROM %s
		WHERE status_id = 2
		ORDER BY purchases DESC, views DESC
	`, coursesTable)

	err := r.db.Select(&courses, query)
	if err != nil {
		return courses, err
	}

	return courses, err
}

func (r *CoursesRepository) GetCourseByID(id int) (model.Course, error) {
	course := model.Course{}

	query := fmt.Sprintf(`
		SELECT 
			id,
			status_id,
			mentor_id,
			category_id,
			name,
			title,
			description,
			effect,
			included_data,
			price,
			views,
			purchases,
			created_at
		FROM %s
		WHERE id = $1
	`, coursesTable)

	if err := r.db.Get(&course, query, id); err != nil {
		return course, err
	}

	return course, nil
}

func NewCoursesRepository(db *sqlx.DB) *CoursesRepository {
	return &CoursesRepository{db: db}
}
