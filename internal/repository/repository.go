package repository

import (
	"github.com/eydeveloper/highload-social/internal/entity"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user entity.User) (string, error)
	GetUser(id string, password string) (entity.User, error)
}

type User interface {
	GetById(id string) (UserProfile, error)
	Search(firstName string, lastName string) ([]UserProfile, error)
}

type Post interface {
	Create(userId string, post entity.Post) (string, error)
	Update(userId string, postId string, post entity.Post) error
	Get(id string) (entity.Post, error)
	Delete(userId string, postId string) error
	GetByIds(postsIds []string) ([]entity.Post, error)
}

type Following interface {
	Follow(followeeId string, followerId string) error
	Unfollow(followeeId string, followerId string) error
	GetFollowers(followeeId string) ([]string, error)
}

type Repository struct {
	Authorization
	User
	Post
	Following
}

func NewRepository(db *sqlx.DB, dbSlave *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(dbSlave),
		Post:          NewPostPostgres(db),
		Following:     NewFollowingPostgres(db),
	}
}
