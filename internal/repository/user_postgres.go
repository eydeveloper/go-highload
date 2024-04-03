package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

type UserProfile struct {
	FirstName string `json:"first_name" db:"first_name" binding:"required"`
	LastName  string `json:"last_name" db:"last_name" binding:"required"`
	BirthDate string `json:"birth_date" db:"birth_date" binding:"required"`
	Gender    string `json:"gender" db:"gender" binding:"required"`
	Biography string `json:"biography" db:"biography"`
	City      string `json:"city" db:"city"`
}

func (r *UserPostgres) GetById(id string) (UserProfile, error) {
	var user UserProfile

	query := fmt.Sprintf(`SELECT first_name, last_name, birth_date, gender, biography, city FROM %s WHERE id = $1`, usersTable)
	err := r.db.Get(&user, query, id)

	return user, err
}

func (r *UserPostgres) Search(firstName string, lastName string) ([]UserProfile, error) {
	var users []UserProfile

	query := fmt.Sprintf(`SELECT first_name, last_name, birth_date, gender, biography, city FROM %s WHERE first_name ILIKE $1 AND last_name ILIKE $2`, usersTable)
	err := r.db.Select(&users, query, firstName+"%", lastName+"%")

	return users, err
}
