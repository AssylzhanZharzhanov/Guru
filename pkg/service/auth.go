package service

import (
	"gitlab.com/zharzhanov/myguru/model"
	"gitlab.com/zharzhanov/myguru/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func (s *AuthService) CreateUser(user model.User) (string, error) {
	return s.repo.CreateUser(user)
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}



