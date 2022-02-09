package service

import "gitlab.com/zharzhanov/myguru/pkg/repository"

type AuthService struct {
	repo repository.Authorization
}

func (a *AuthService) CreateUser() (string, error) {
	panic("implement me")
}

func (a *AuthService) GetUser() (string, error) {
	panic("implement me")
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}



