package seeds

import (
	"github.com/eydeveloper/highload-social/internal/entity"

	"github.com/go-faker/faker/v4"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UsersSeeder struct {
	DB    *sqlx.DB
	Count int
}

func NewUsersSeeder(db *sqlx.DB, count int) *UsersSeeder {
	return &UsersSeeder{DB: db, Count: count}
}

func (u *UsersSeeder) Seed() error {
	tx, err := u.DB.Beginx()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.Preparex(`INSERT INTO users (first_name, last_name, birth_date, password_hash, gender, biography, city) VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return err
	}

	defer stmt.Close()

	for i := 0; i < u.Count; i++ {
		user := entity.User{}

		err := faker.FakeData(&user)
		if err != nil {
			return err
		}

		_, err = stmt.Exec(user.FirstName, user.LastName, user.BirthDate, user.Password, user.Gender, user.Biography, user.City)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
