package service

import "gitlab.com/zharzhanov/myguru/pkg/repository"

type Authorization interface {

}

type Service struct {
	Authorization
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
	}
}
