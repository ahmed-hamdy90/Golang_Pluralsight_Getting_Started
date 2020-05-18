package main

import (
	"fmt"
)

func main() {
	// use simple declare integer variable then re-sign it
	var i int
	i = 42
	fmt.Println(i)

	// use simple declare float variable with assign value
	var f float32 = 3.14
	fmt.Println(f)

	// use implicit type to short assignment variable
	firstName := "Arthur"
	fmt.Println(firstName)

	// use simple declare boolean variable with assign value
	b := true
	fmt.Println(b)

	// use complex built-in method
	c := complex(3, 4)
	fmt.Println(c)
	// separate complex number to real and imaginary numbers
	// use implicit type to assign two variables on the same line
	r, im := real(c), imag(c)
	fmt.Println(r, im)
}