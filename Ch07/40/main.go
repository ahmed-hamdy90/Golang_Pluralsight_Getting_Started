package main

func main() {
	// Create slice
	slice := []int {1, 2, 3}

	// Begin loop on slice (loop on collections type) => the ugly syntax
	for i := 0; i < len(slice); i++ {
		// Use println built-in method
		println(slice[i])
	}

	// Begin loop on slice (loop on collections type) => the clean Go syntax
	// use built-in range keyword which return index/key and value of any collections
	//for i, v := range slice {
	//	println(i, v)
	//}

	// Create map
	//wellKnownPorts := map[string]int {"http": 80, "https": 443}

	// Begin loop on map (loop on collections type) => the clean Go syntax
	//for k, v := range wellKnownPorts {
	//	println(k, v)
	//}

	// Begin loop on keys of map only (loop on collections type) => the clean Go syntax
	//for k := range wellKnownPorts {
	//	println(k)
	//}

	// Begin loop on values of map only (loop on collections type) => the clean Go syntax
	//for _, v := range wellKnownPorts {
	//	println(v)
	//}
}