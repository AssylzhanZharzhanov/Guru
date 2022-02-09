package service

import "gitlab.com/zharzhanov/myguru/pkg/repository"

type Authorization interface {
	CreateUser() (string, error)
	GetUser() (string, error)
}

type Courses interface {
	GetCourses()
}

type Service struct {
	Authorization
	Courses
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		Courses: NewCoursesService(repository.Courses),
	}
}
