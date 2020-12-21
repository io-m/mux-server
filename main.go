package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	h "github.com/io-m/mux-server/handler"
	"github.com/io-m/mux-server/model"
)

func main() {
	// Defining new instance of Gorilla mux
	m := mux.NewRouter()
	// ==================
	// Defining handlers
	// m.HandleFunc("/form", h.ViewForm).Methods("GET")
	m.HandleFunc("/register", h.RegisterForm).Methods("GET")
	m.HandleFunc("/login", h.LoginForm).Methods("GET")
	m.HandleFunc("/register", h.Register).Methods("POST")
	m.HandleFunc("/login", h.Login).Methods("POST")
	m.HandleFunc("/people", h.GetAll).Methods("GET")
	m.HandleFunc("/people/{personID}", h.GetOne).Methods("GET")
	m.HandleFunc("/people", h.Create).Methods("POST")
	m.HandleFunc("/people/{personID}", h.Update).Methods("PATCH")
	m.HandleFunc("/people/{personID}", h.Destroy).Methods("DELETE")

	// ======================================
	log.Fatal(http.ListenAndServe(":7700", m))
	defer model.Db.Close()

}
