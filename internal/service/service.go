package service

import (
	"github.com/eydeveloper/highload-social/internal/entity"
	"github.com/eydeveloper/highload-social/internal/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (string, error)
	GenerateToken(id string, password string) (string, error)
	ParseToken(token string) (string, error)
}

type User interface {
	GetById(id string) (repository.UserProfile, error)
	Search(firstName string, lastName string) ([]repository.UserProfile, error)
}

type Post interface {
	Create(userId string, post entity.Post) (string, error)
	Update(userId string, postId string, post entity.Post) error
	Get (id string) (entity.Post, error)
	Delete (userId string, postId string) error
}

type Following interface {
	Follow(followeeId string, followerId string) error
	Unfollow(followeeId string, followerId string) error
}

type Service struct {
	Authorization
	User
	Post
	Following
}

func NewService(repositories *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repositories.Authorization),
		User:          NewUserService(repositories.User),
		Post:          NewPostService(repositories.Post),
		Following:     NewFollowingService(repositories.Following),
	}
}
