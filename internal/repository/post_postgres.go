package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/eydeveloper/highload-social/internal/entity"
	"github.com/jmoiron/sqlx"
)

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

func (r *PostPostgres) Create(userId string, post entity.Post) (string, error) {
	var id string

	query := fmt.Sprintf(`INSERT INTO %s (author_id, content) VALUES ($1, $2) RETURNING id`, userPostsTable)
	row := r.db.QueryRow(query, userId, post.Content)

	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *PostPostgres) Update(userId string, postId string, post entity.Post) error {
	var id string

	query := fmt.Sprintf("SELECT id FROM %s WHERE id = $1 AND author_id = $2 LIMIT 1", userPostsTable)
	err := r.db.Get(&id, query, postId, userId)

	if errors.Is(err, sql.ErrNoRows) {
		return err
	}

	query = fmt.Sprintf("UPDATE user_posts SET content = $1 WHERE id = $2")
	_, err = r.db.Exec(query, post.Content, postId)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostPostgres) Get(id string) (entity.Post, error) {
	var post entity.Post

	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1 LIMIT 1`, userPostsTable)
	err := r.db.Get(&post, query, id)

	return post, err
}

func (r *PostPostgres) Delete(userId string, postId string) error {
	var id string

	query := fmt.Sprintf("SELECT FROM %s WHERE id = $1 AND author_id = $2 LIMIT 1", userPostsTable)
	err := r.db.Get(&id, query, postId, userId)

	if errors.Is(err, sql.ErrNoRows) {
		return err
	}

	query = fmt.Sprintf("DELETE FROM %s WHERE id = $1", userPostsTable)
	_, err = r.db.Exec(query, postId)

	if err != nil {
		return err
	}

	return nil
}

func (r *PostPostgres) GetByIds(postsIds []string) ([]entity.Post, error) {
	var posts []entity.Post

	query := fmt.Sprintf("SELECT id, author_id, content, created_at FROM %s WHERE id IN (?)", userPostsTable)
	query, args, err := sqlx.In(query, postsIds)

	if err != nil {
		return nil, err
	}

	query = r.db.Rebind(query)
	err = r.db.Select(&posts, query, args...)

	if err != nil {
		return nil, err
	}

	return posts, nil
}
