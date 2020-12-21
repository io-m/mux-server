package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/io-m/mux-server/model"
	"github.com/io-m/mux-server/utils"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("views/*html"))
}

// LoginForm displays login form as template
func LoginForm(w http.ResponseWriter, r *http.Request) {
	// Querying data in DB and send it to FetchedUser instance

	// rows, err := model.Db.Query("SELECT * FROM people.user")
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	// var id int
	// var name string
	// var eMail string
	// var passw string
	// // Users is lice of all users from DB
	// var Users []model.User
	// for rows.Next() {
	// 	if err := rows.Scan(&id, &name, &eMail, &passw); err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		return
	// 	}
	// 	oneUser := model.User{ID: id, UserName: name, Email: eMail, Password: passw}
	// 	Users = append(Users, oneUser)
	// }

	user := utils.UserModel{
		DB: model.Db,
	}
	User, err := user.FetchOne(Email)
	if err != nil {
		log.Fatal(err)
	}

	tpl.ExecuteTemplate(w, "login.html", User)
}

// func ViewForm(w http.ResponseWriter, r *http.Request) {
// 	tpl.ExecuteTemplate(w, "form.html", nil)
// }

// RegisterForm displays registration form
func RegisterForm(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "register.html", nil)
}
