package main

import (
	"fmt"
)

func main() {
	// define new Struct type
	type user struct {
		ID int
		FirstName string
		LastName string
	}

	// create new instance from user struct type
	var u user // empty instance

	fmt.Println(u)

	// begin assign members under initialize user struct type
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Dent"

	fmt.Println(u)

	// create new instance from user struct type with short initialize using implicit type
	u2 := user {ID: 1, FirstName: "Arthur", LastName: "Dent"}

	fmt.Println(u2)
}