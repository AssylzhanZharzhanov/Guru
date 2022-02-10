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

func (s *CoursesService) GetTrendingCourses() ([]model.Course, error) {
	return s.repo.GetTrendingCourses()
}

func (s *CoursesService) GetComingSoonCourses() ([]model.Course, error) {
	return s.repo.GetComingSoonCourses()
}

func (s *CoursesService) GetCourseByID(id int) (model.Course, error) {
	return s.repo.GetCourseByID(id)
}

func NewCoursesService(repo repository.Courses) *CoursesService {
	return &CoursesService{repo: repo}
}
