package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	// call custom new method
	startWebServer()
}

// create custom new method rather than default main method
func startWebServer() {
	fmt.Println("Starting server...")
	// do important thing
	fmt.Println("Server started")
}