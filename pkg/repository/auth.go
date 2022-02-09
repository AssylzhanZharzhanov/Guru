package repository

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/zharzhanov/myguru/model"
)

type AuthRepository struct {
	db *sqlx.DB
}

func (r *AuthRepository) CreateUser(user model.User) (string, error) {
	panic("implement me")
}

func (r *AuthRepository) GetUser(user model.User) (string, error) {
	panic("implement me")
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}