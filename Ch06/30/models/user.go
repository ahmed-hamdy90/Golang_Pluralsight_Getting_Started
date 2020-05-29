package models

// Create User model with Struct type
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// Declare global variables using var block
var (
	users  []*User // slice to save pointers of created users
	nextID = 1     // referred as insert unique ID (look like Database)
)

// Create custom method to retrieve exists users slice
func GetUsers() []*User {
	return users
}

// Create custom method to add new User to exists users slice
func AddUser(u User) (User, error) {
	// re-set user's id with nextID counter member which act as insert into Database
	u.ID = nextID
	// then increment this counter member
	nextID++
	// add pointer of given user under exists users slice using built-in append method
	users = append(users, &u)

	return u, nil
}