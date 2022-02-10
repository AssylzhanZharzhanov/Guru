package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/zharzhanov/myguru/model"
)

type CategoriesRepository struct {
	db *sqlx.DB
}

func (r *CategoriesRepository) GetCategories(lang string) ([]model.Category, error) {
	categories := make([]model.Category, 0)

	query := fmt.Sprintf(`
		SELECT 
			id,
			name
		FROM %s
		ORDER BY id ASC
	`, categoriesTable)

	err := r.db.Select(&categories, query)
	if err != nil {
		return categories, err
	}

	return categories, err
}

func NewCategoriesRepository(db *sqlx.DB) *CategoriesRepository {
	return &CategoriesRepository{db: db}
}

