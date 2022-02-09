package service

import (
	"gitlab.com/zharzhanov/myguru/model"
	"gitlab.com/zharzhanov/myguru/pkg/repository"
)

type MentorsService struct {
	repo repository.Mentors
}

func (m MentorsService) GetMentors() ([]model.Mentor, error) {
	panic("implement me")
}

func (m MentorsService) GetMentorByID() (model.Mentor, error) {
	panic("implement me")
}

func NewMentorsService(repo repository.Mentors) *MentorsService {
	return &MentorsService{repo: repo}
}