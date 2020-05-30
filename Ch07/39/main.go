package main

func main() {
	var i int

	// Use infinite loop => the ugly syntax
	for ; ;  {
		if i == 5 {
			break
		}
		// use println built-in method
		println(i)
		i++
	}

	// Use infinite loop => the clean Go syntax
	//for {
	//	if i == 5 {
	//		break
	//	}
	//	println(i)
	//	i++
	//}
}