package entity

type Post struct {
	Id        string `json:"id" db:"id"`
	AuthorId  string `json:"author_id" db:"author_id"`
	Content   string `json:"content" db:"content" binding:"required"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
