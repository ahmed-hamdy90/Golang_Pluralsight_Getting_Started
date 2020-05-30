package models

import (
	"errors"
	"fmt"
)

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
	// make sure given User instance must be has not id value
	if u.ID != 0 {
		// If happened will return empty User instance and error as Return value foe method
		return User{}, errors.New("New User must not include id or it must be set to zero")
	}
	// re-set user's id with nextID counter member which act as insert into Database
	u.ID = nextID
	// then increment this counter member
	nextID++
	// add pointer of given user under exists users slice using built-in append method
	users = append(users, &u)

	return u, nil
}

func GetUserByID(id int) (User, error) {
	// loop on exists users to fetch given id
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	// if not found will return empty User instance and error => will use Errorf method under fmt library which return error instance
	return User{}, fmt.Errorf("User with Id '%v' not found", id)
}

func UpdateUser(u User) (User, error) {
	// loop on exists users to fetch old instance of given User
	for i, candidate := range users {
		if candidate.ID == u.ID {
			// replace the old instance with new one, then return it
			users[i] = &u

			return u, nil
		}
	}

	// if not found will return empty User instance and error
	return User{}, fmt.Errorf("User with Id '%v' not found", u.ID)
}

func RemoveUserById(id int) error {
	// loop on exists users to fetch given id
	for i, u := range users {
		if u.ID == id {
			// remove user will process using append built-in method => append items before and after the removed item
			users = append(users[:i], users[i+1:]...)

			return nil
		}
	}

	// if not found will error
	return fmt.Errorf("User with Id '%v' not found", id)
}