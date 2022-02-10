package service

import (
	"gitlab.com/zharzhanov/myguru/model"
	"gitlab.com/zharzhanov/myguru/pkg/repository"
)

type CoursesService struct {
	repo repository.Courses
}

func (s *CoursesService) GetCourses(lang string) ([]model.Course, error) {
	return s.repo.GetCourses(lang)
}

func (s *CoursesService) GetTrendingCourses(lang string) ([]model.Course, error) {
	return s.repo.GetTrendingCourses(lang)
}

func (s *CoursesService) GetComingSoonCourses(lang string) ([]model.Course, error) {
	return s.repo.GetComingSoonCourses(lang)
}

func (s *CoursesService) GetCourseByID(id int, lang string) (model.Course, error) {
	return s.repo.GetCourseByID(id, lang)
}

func NewCoursesService(repo repository.Courses) *CoursesService {
	return &CoursesService{repo: repo}
}
