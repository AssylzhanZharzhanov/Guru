package repository

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/zharzhanov/myguru/model"
)

type MentorsRepository struct {
	db *sqlx.DB	
}

func (m MentorsRepository) GetMentors() ([]model.Mentor, error) {
	panic("implement me")
}

func (m MentorsRepository) GetMentorByID() (model.Mentor, error) {
	panic("implement me")
}

func NewMentorsRepository(db *sqlx.DB) *MentorsRepository {
	return &MentorsRepository{db: db}
}
