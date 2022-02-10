package service

import (
	"gitlab.com/zharzhanov/myguru/model"
	"gitlab.com/zharzhanov/myguru/pkg/repository"
)

type CategoriesService struct {
	repo repository.Categories
}

func (s *CategoriesService) GetCategories(lang string) ([]model.Category, error) {
	return s.repo.GetCategories(lang)
}

func NewCategoriesService(repository repository.Categories) *CategoriesService {
	return &CategoriesService{}
}

