package service

import (
	"github.com/eydeveloper/highload-social/internal/entity"
	"github.com/eydeveloper/highload-social/internal/repository"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) Create(userId string, post entity.Post) (string, error) {
	return s.repo.Create(userId, post)
}

func (s *PostService) Update(userId string, postId string, post entity.Post) error {
	return s.repo.Update(userId, postId, post)
}

func (s *PostService) Get(postId string) (entity.Post, error) {
	return s.repo.Get(postId)
}

func (s *PostService) Delete(userId string, postId string) error {
	return s.repo.Delete(userId, postId)
}
