package service

import (
	"gitlab.com/zharzhanov/myguru/model"
	"gitlab.com/zharzhanov/myguru/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (string, error)
}

type Categories interface {
	GetCategories(lang string) ([]model.Category, error)
}

type Courses interface {
	GetCourses(lang string) ([]model.Course, error)
	GetTrendingCourses(lang string) ([]model.Course, error)
	GetComingSoonCourses(lang string) ([]model.Course, error)
	GetCourseByID(id int, lang string) (model.Course, error)
}

type Mentors interface {
	GetMentors() ([]model.Mentor, error)
	GetMentorByID(id int) (model.Mentor, error)
}

type Service struct {
	Authorization
	Categories
	Courses
	Mentors
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		Categories: NewCategoriesService(repository.Categories),
		Courses: NewCoursesService(repository.Courses),
		Mentors: NewMentorsService(repository.Mentors),
	}
}
