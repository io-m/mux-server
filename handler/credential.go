package handler

import (
	"log"
	"net/http"

	"github.com/io-m/mux-server/model"
)

// Email is used as unique key for fetching single item from DB
var Email string

// Register is handler for registering user and pushing it to DB
func Register(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	username := r.FormValue("username")
	Email = r.FormValue("email")
	pass := r.FormValue("password")

	// FetchedUser is user model instance populated with data from DB
	// FetchedUser := model.User{
	// 	ID:       id,
	// 	UserName: username,
	// 	Email:    email,
	// 	Password: pass,
	// }

	// if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
	// 	log.Fatal(err)
	// }
	inCommand, err := model.Db.Prepare("INSERT INTO people.user (username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer inCommand.Close()
	_, err = inCommand.Exec(username, Email, pass)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
	}
	log.Printf("Successfully logged in\n")
	// =======================================

	http.Redirect(w, r, "/login", http.StatusFound)
}

// Login is handler for loging user in and checking credentials against DB
func Login(w http.ResponseWriter, r *http.Request) {

}
