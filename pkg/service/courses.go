package service

import (
	"gitlab.com/zharzhanov/myguru/model"
	"gitlab.com/zharzhanov/myguru/pkg/repository"
)

type CoursesService struct {
	repo repository.Courses
}

func (s *CoursesService) GetCourses() ([]model.Course, error) {
	return s.repo.GetCourses()
}

func (s *CoursesService) GetCourseByID() (model.Course, error) {
	panic("implement me")
}

func NewCoursesService(repo repository.Courses) *CoursesService {
	return &CoursesService{repo: repo}
}
