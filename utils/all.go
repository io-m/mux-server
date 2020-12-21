package utils

import (
	"database/sql"

	"github.com/io-m/mux-server/model"
)

// UserModel is struct that represent User data
type UserModel struct {
	DB *sql.DB
}

// FetchAll is method of fetching all data from DB
func (um *UserModel) FetchAll() ([]model.User, error) {
	var id int
	var username string
	var email string
	var password string

	rows, err := um.DB.Query("SELECT * FROM user")
	if err != nil {
		return nil, err
	}
	Users := []model.User{}
	for rows.Next() {
		if err = rows.Scan(&id, &username, &email, &password); err != nil {
			return nil, err
		}
		user := model.User{ID: id, UserName: username, Email: email, Password: password}
		Users = append(Users, user)
	}
	return Users, nil
}
