package service

import (
	"github.com/eydeveloper/highload-social/internal/entity"
	"github.com/eydeveloper/highload-social/internal/repository"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
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
	Create(userId string, post entity.Post) (entity.Post, error)
	Update(userId string, postId string, post entity.Post) error
	Get(id string) (entity.Post, error)
	Delete(userId string, postId string) error
}

type Feed interface {
	Get(userId string) ([]entity.Post, error)
	GetRealTime(userId string) (<-chan amqp.Delivery, error)
	AddPost(userId string, post entity.Post) error
}

type Following interface {
	Follow(followeeId string, followerId string) error
	Unfollow(followeeId string, followerId string) error
	GetFollowers(followeeId string) ([]string, error)
}

type Service struct {
	Authorization
	User
	Post
	Feed
	Following
}

func NewService(repositories *repository.Repository, redisClient *redis.Client, amqpChannel *amqp.Channel) *Service {
	return &Service{
		Authorization: NewAuthService(repositories.Authorization),
		User:          NewUserService(repositories.User),
		Post:          NewPostService(repositories.Post),
		Feed:          NewFeedService(repositories.Post, redisClient, amqpChannel),
		Following:     NewFollowingService(repositories.Following),
	}
}
