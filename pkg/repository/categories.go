package repository

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/zharzhanov/myguru/model"
)

type CategoriesRepository struct {
	db *sqlx.DB
}

func (r *CategoriesRepository) GetCategories(lang string) ([]model.Category, error) {
	panic("implement me")
}

func NewCategoriesRepository(db *sqlx.DB) *CategoriesRepository {
	return &CategoriesRepository{db: db}
}

