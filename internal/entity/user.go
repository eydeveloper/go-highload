package entity

type User struct {
	Id        string `json:"-" db:"id"`
	FirstName string `json:"first_name" db:"first_name" binding:"required" faker:"first_name"`
	LastName  string `json:"last_name" db:"last_name" binding:"required" faker:"last_name"`
	BirthDate string `json:"birth_date" db:"birth_date" binding:"required" faker:"date"`
	Gender    string `json:"gender" db:"gender" binding:"required" faker:"oneof:M,F"`
	Biography string `json:"biography" db:"biography" faker:"sentence"`
	City      string `json:"city" db:"city" faker:"word"`
	Password  string `json:"password" binding:"required" faker:"password"`
}
