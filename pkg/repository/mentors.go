package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/zharzhanov/myguru/model"
)

type MentorsRepository struct {
	db *sqlx.DB	
}

func (r *MentorsRepository) GetMentors() ([]model.Mentor, error) {
	mentors := make([]model.Mentor, 0)

	query := fmt.Sprintf(`
		SELECT
			id,
			first_name,
			last_name,
			info,
			created_at
		FROM %s
	`, mentorsTable)

	err := r.db.Select(&mentors, query)
	if err != nil {
		return mentors, err
	}

	return mentors, err
}

func (r *MentorsRepository) GetMentorByID(id int) (model.Mentor, error) {
	mentor := model.Mentor{}

	query := fmt.Sprintf(`
		SELECT
			id,
			first_name,
			last_name,
			info,
			created_at
		FROM %s
		WHERE id = $1
	`, mentorsTable)

	if err := r.db.Get(&mentor, query, id); err != nil {
		return mentor, err
	}

	return mentor, nil
}

func NewMentorsRepository(db *sqlx.DB) *MentorsRepository {
	return &MentorsRepository{db: db}
}
