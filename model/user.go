package model

// User struct is model for interacting between client and DB
type User struct {
	ID       int    `json:"id" sql:"id"`
	UserName string `json:"username" sql:"username"`
	Email    string `json:"email" sql:"email"`
	Password string `json:"password" sql:"password"`
}
