package entity

type User struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	UserName   string `json:"user_name"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	UserStatus string `json:"user_status"`
}
