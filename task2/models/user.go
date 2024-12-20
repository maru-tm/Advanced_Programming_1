package models

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" gorm:"unique"`
	Email string `json:"email"`
}
