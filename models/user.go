package models

type User struct {
	Username string `json:"username"`
	Id       string `json:"id"`
	Password string `json:"passoword"`
	Role     string `json:"role"`
}
