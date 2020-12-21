package utils

import "github.com/io-m/mux-server/model"

// FetchOne is method on model that fetches only one specified row
func (um *UserModel) FetchOne(eMail string) (model.User, error) {
	var id int
	var username string
	var email string
	var password string

	rows, err := um.DB.Query("SELECT * FROM user WHERE email = ?", eMail)
	if err != nil {
		return model.User{}, err
	}
	var User model.User
	for rows.Next() {
		if err = rows.Scan(&id, &username, &email, &password); err != nil {
			return model.User{}, err
		}
		User = model.User{ID: id, UserName: username, Email: email, Password: password}
	}
	return User, nil
}
