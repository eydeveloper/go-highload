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

type Repository struct {
	Authorization
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
	}
}
