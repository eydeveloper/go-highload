package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type FollowingPostgres struct {
	db *sqlx.DB
}

func NewFollowingPostgres(db *sqlx.DB) *FollowingPostgres {
	return &FollowingPostgres{db: db}
}

func (r *FollowingPostgres) Follow(followeeId string, followerId string) error {
	var id string

	query := fmt.Sprintf("SELECT FROM %s WHERE user_id = $1 AND follower_id = $2 LIMIT 1", userFollowersTable)
	err := r.db.Get(&id, query, followeeId, followerId)

	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	query = fmt.Sprintf("INSERT INTO %s (user_id, follower_id) VALUES ($1, $2) RETURNING id", userFollowersTable)
	err = r.db.Get(&id, query, followeeId, followerId)

	if err != nil {
		return err
	}

	return nil
}

func (r *FollowingPostgres) Unfollow(followeeId string, followerId string) error {
	var id string

	query := fmt.Sprintf("SELECT id FROM %s WHERE user_id = $1 AND follower_id = $2 LIMIT 1", userFollowersTable)
	err := r.db.Get(&id, query, followeeId, followerId)

	if errors.Is(err, sql.ErrNoRows) {
		return err
	}

	query = fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND follower_id = $2", userFollowersTable)
	_, err = r.db.Exec(query, followeeId, followerId)

	if err != nil {
		return err
	}

	return nil
}

func (r *FollowingPostgres) GetFollowers(followeeId string) ([]string, error) {
	var followers []string

	query := fmt.Sprintf("SELECT follower_id FROM %s WHERE user_id = $1", userFollowersTable)
	err := r.db.Select(&followers, query, followeeId)

	if err != nil {
		return nil, err
	}

	return followers, nil
}
