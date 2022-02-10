package service

import (
	"gitlab.com/zharzhanov/myguru/model"
	"gitlab.com/zharzhanov/myguru/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (string, error)
}

type Courses interface {
	GetCourses() ([]model.Course, error)
	GetTrendingCourses() ([]model.Course, error)
	GetComingSoonCourses() ([]model.Course, error)
	GetCourseByID(id int) (model.Course, error)
}

type Mentors interface {
	GetMentors() ([]model.Mentor, error)
	GetMentorByID(id int) (model.Mentor, error)
}

type Service struct {
	Authorization
	Courses
	Mentors
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		Courses: NewCoursesService(repository.Courses),
		Mentors: NewMentorsService(repository.Mentors),
	}
}
