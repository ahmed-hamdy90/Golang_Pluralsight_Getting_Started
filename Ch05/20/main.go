package main

import (
	"fmt"
)

func main() {

	// Create an array with predefine size
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)
	fmt.Println(arr[2])

	//fmt.Println(arr[4]) // will throw compiler error as array with 3 elements

	// Create an array using implicit type also include values under array too
	//arr2 := [3]int {1, 2, 3}
	//
	//fmt.Println(arr2)
}