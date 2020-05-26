package main

import (
	"fmt"
)

func main() {
	// create an array with fixed 3 element size using implicit type
	arr := [3]int {1, 2, 3}

	// create slice from created array using implicit type [with : column operator which take all element from an array]
	slice := arr[:]

	fmt.Println(arr, slice)

	// re-set element under create array and slice, will reflected on both them as they point to same elements position (call by reference)
	arr[1] = 42
	slice[2] = 27

	fmt.Println(arr, slice)

	// create slice using implicit type and underline array(array with not fix size)
	//slice := []int {1, 2, 3}
	//
	//fmt.Println(slice)

	// use built-in append method to add new element to array
	//slice = append(slice, 4, 42, 27)
	//
	//fmt.Println(slice)

	// use column operator to create new slices from first created slice
	//s2 := slice[1:] // take elements under created slice from the index 1 -> to the end
	//s3 := slice[:2] // take elements under created slice from the beginning -> to the index 2
	//s4 := slice[1:2] // take elements under created slice from the index 1 -> to the index 2
	//
	//fmt.Println(s2, s3, s4)
}