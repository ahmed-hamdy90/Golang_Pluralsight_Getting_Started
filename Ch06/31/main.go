package main

import (
	"fmt"
	"github.com/pluralsight/webservice/models"
)

func main() {
	// initialize user instance using implicit type
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}

	fmt.Println(u)
}
