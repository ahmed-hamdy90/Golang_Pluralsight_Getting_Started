package models

// Create User model using Struct type
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// Declare global variables using var block
var (
	users  []*User // slice to save pointers of created users instance
	nextID = 1     // referred as insert unique ID (look like Database)
)
