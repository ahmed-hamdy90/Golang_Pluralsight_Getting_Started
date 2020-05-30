package main

// Create User struct type
type User struct {
	ID int
	FirstName string
	LastName string
}

func main() {
	// create two instance form User
	u1 := User {
		ID: 1,
		FirstName: "Arthur",
		LastName: "Dent",
	}

	u2 := User {
		ID: 2,
		FirstName: "Fred",
		LastName: "Prefect",
	}

	// Use IF statement to compare two User instance
	if u1.ID == u2.ID {
		// use built-in println method
		println("Same user!")
	}

	//if u1.ID != u2.ID {
	//	println("Not same user!")
	//}

	// Use IF, else statement to compare two User instance
	//if u1.ID == u2.ID {
	//	println("Same user!")
	//} else {
	//	println("Different user!")
	//}

	// Use IF, else IF, else statement to compare two User instance
	//if u1.ID == u2.ID {
	//	println("Same user!")
	//} else if u1.FirstName == u2.FirstName {
	//	println("Similar user!")
	//} else {
	//	println("Different user!")
	//}
}