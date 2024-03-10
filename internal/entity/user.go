package entity

type User struct {
	Id        string `json:"-" db:"id"`
	FirstName string `json:"first_name" db:"first_name" binding:"required"`
	LastName  string `json:"last_name" db:"last_name" binding:"required"`
	BirthDate string `json:"birth_date" db:"birth_date" binding:"required"`
	Gender    string `json:"gender" db:"gender" binding:"required"`
	Biography string `json:"biography" db:"biography"`
	City      string `json:"city" db:"city"`
	Password  string `json:"password" binding:"required"`
}
