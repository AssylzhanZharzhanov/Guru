package service

import "gitlab.com/zharzhanov/myguru/pkg/repository"

type CoursesService struct {
	repo repository.Courses
}

func (s *CoursesService) GetCourses() {
	s.repo.GetCourses()
}

func NewCoursesService(repo repository.Courses) *CoursesService {
	return &CoursesService{repo: repo}
}
