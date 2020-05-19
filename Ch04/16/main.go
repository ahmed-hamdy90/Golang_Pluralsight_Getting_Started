package main

import (
	"fmt"
)

func main() {
	// create a constant with assign value
	const pi = 3.1414
	fmt.Println(pi)

	// try re-assign constant
	// pi = 1.2 // will throw error

	// create a constant with implicit type and used it into operation
	// const c = 3
	// fmt.Println(c + 3)
	//
	// // a bunch of code
	//
	// fmt.Println(c + 1.2)

	// create a constant with restrict type(explicit type) and used it into operation
	// const c int = 3
	// fmt.Println(c + 3)
	//
	// // a bunch of code
	//
	// fmt.Println(c + 1.2) // this line with throw error as can't add integer to float member type
	//
	// // fix above error using float32 convert method
	// fmt.Println(float32(c) + 1.2)
}