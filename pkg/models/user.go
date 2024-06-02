package models

type User struct {
	ID           string `json:"id"`
	Fullname     string `json:"fullname"`
	Email        string `json:"email"`
	HashPassword string `json:"hashPassword"`
	Books        []Book `json:"books"`
}
