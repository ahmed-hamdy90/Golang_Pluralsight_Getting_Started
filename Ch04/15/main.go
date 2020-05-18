package main

import(
	"fmt"
)

func main() {
	// create pointer without any assignment
	// var firstName *string
	// fmt.Println(firstName)

	// assign value for declare pointer using Dereference => given error as already this point hasn't space reserved
	// *firstName = "Arther"
	// fmt.Println(firstName)

	// create pointer with refer to string and reserve space
	// var firstName *string = new(string)
	// *firstName = "Arthur"
	// fmt.Println(firstName) // print pointer id not string value which pointer point to
	// fmt.Println(*firstName) // print string value which pointer point to

	// create variable
	firstName := "Arthur"

	// create pointer point on this variable
	ptr := &firstName
	fmt.Println(ptr, *ptr) // print both pointer id and string value which pointer point to

	// re-assign variable
	firstName = "Tricia"
	fmt.Println(ptr, *ptr) // re-print both pointer id and string value which pointer point to
}