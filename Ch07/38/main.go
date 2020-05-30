package main

func main() {
	var i int

	// Use loop till condition with a post clause
	for ; i < 5; i++ {
		// use println built-in method
		println(i)
	}

	// Use loop till condition with a post clause include declare i variable
	//for i := 0; i < 5; i++ {
	//	println(i)
	//}

	// For the first loop will print the latest i variable's value
	// For the second loop will throw error as i variable under loop's scope only (we need to comment it to work)
	println(i)
}