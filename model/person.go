// Package model defines basic Person model for interacting between client and database
package model

// Person is model struct for DB people
type Person struct {
	ID        int    `json:"id" sql:"id"`
	FirstName string `json:"firstname" sql:"fName"`
	LastName  string `json:"lastname" sql:"lName"`
}
