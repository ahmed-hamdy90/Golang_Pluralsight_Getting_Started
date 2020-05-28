package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")

	// create new variable for web server's port using implicit type
	port := 3000

	// call custom new method with passing port variable as parameter
	startWebServer(port)

	// call custom new method with passing port variable and number of retries as parameters
	//startWebServer(port, 2)
}

// create custom Web server method with given port as parameter
func startWebServer(port int) {
	fmt.Println("Starting server...")
	// do important thing
	fmt.Println("Server started on port", port)
}

// Create custom Web server method with two given parameters: port for server and number of retries
//func startWebServer(port int, numberOfRetries int) {
//	fmt.Println("Starting server...")
//	// do important thing
//	fmt.Println("Server started on port", port)
//	fmt.Println("Number of retries", numberOfRetries)
//}

// The same of above method, But change way of method's parameters declaration to define parameter's type once
// As both parameters with same type so can combine them like blow
//func startWebServer(port, numberOfRetries int) {
//	fmt.Println("Starting server...")
//	// do important thing
//	fmt.Println("Server started on port", port)
//	fmt.Println("Number of retries", numberOfRetries)
//}
