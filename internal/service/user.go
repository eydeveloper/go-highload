package service

import (
	"github.com/eydeveloper/highload-social/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetById(id string) (repository.UserProfile, error) {
	return s.repo.GetById(id)
}

func (s *UserService) Search(firstName string, lastName string) ([]repository.UserProfile, error) {
	return s.repo.Search(firstName, lastName)
}
