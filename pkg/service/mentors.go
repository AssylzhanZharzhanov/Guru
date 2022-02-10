package service

import (
	"gitlab.com/zharzhanov/myguru/model"
	"gitlab.com/zharzhanov/myguru/pkg/repository"
)

type MentorsService struct {
	repo repository.Mentors
}

func (s *MentorsService) GetMentors() ([]model.Mentor, error) {
	return s.repo.GetMentors()
}

func (s *MentorsService) GetMentorByID(id int) (model.Mentor, error) {
	return s.repo.GetMentorByID(id)
}

func NewMentorsService(repo repository.Mentors) *MentorsService {
	return &MentorsService{repo: repo}
}