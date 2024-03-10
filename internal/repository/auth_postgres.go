package repository

import (
	"fmt"

	"github.com/eydeveloper/highload-social/internal/entity"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user entity.User) (string, error) {
	var id string

	query := fmt.Sprintf(`
		INSERT INTO %s
		(
			first_name,
			last_name,
			birth_date,
			gender,
			biography,
			city,
			password_hash
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`, usersTable)

	row := r.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.BirthDate,
		user.Gender,
		user.Biography,
		user.City,
		user.Password,
	)

	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(id string, password string) (entity.User, error) {
	var user entity.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE id=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, id, password)

	return user, err
}
