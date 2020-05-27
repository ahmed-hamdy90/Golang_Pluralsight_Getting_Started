package models

// create User model using Struct type
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// declare global variables using var block
var (
	users  []*User // slice to save users instance collections
	nextID = 1     // referred as insert unique ID (look line Database)
)
