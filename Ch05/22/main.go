package main

import (
	"fmt"
)

func main() {
	// create map using implicit type and short declare syntax map[key-type]value-type
	m := map[string]int {"foo": 42}

	fmt.Println(m)
	fmt.Println(m["foo"])

	// re-set one element's value under map
	m["foo"] = 27

	fmt.Println(m)

	// delete an element from map using delete built-in method
	delete(m, "foo")

	fmt.Println(m)
}