package main

import (
	"fmt"
)

// Create the block of enumerated constants (constant block)
const (
	first = 1
	second = "second"
)

// Create the block of enumerated constants and use iot (an enumerator for constant block)
// by default iot begin with 0 value and increment by 1
//const (
//	first = iota
//	second = iota
//)

// Another usage for iot (use iot with mathematic operation)
//const (
//	first = iota + 6
//	second = 2 << iota
//)

// Can declare constant without setting it's value, value will be taken from previous declare constant increment by 1
//const (
//	first = iota + 6
//	second
//)

// Can declare two (more then one) blocks of enumerated constants,
// plus use iot into two constants blocks, the behaviour of iot will rest every constant block with beginning value (zero)
//const (
//	first = iota
//	second
//)
//
//const (
//	third = iota
//)

// Another example for declare two blocks of enumerated constants
//const (
//	first = iota
//	second = iota
//)
//
//const (
//	third = iota
//	fourth
//)

func main() {
	fmt.Println(first, second)

	//fmt.Println(first, second, third) // in case example number 5 only

	//fmt.Println(first, second, third, fourth) // in case example number 6 (the last one) only
}