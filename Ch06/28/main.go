package main

import (
	"fmt"
)

func main() {
	// create new variable for web server's port using implicit type
	port := 3000

	// call custom new method with passing port variable as parameter, get returned Boolean value from method
	isStarted := startWebServer(port)
	fmt.Println(isStarted)

	// call custom new method with passing port variable as parameter, get returned error from method
	//err := startWebServer(port)
	//fmt.Println(err)

	// call custom new method with passing port variable as parameter
	// get multiple returned values integer and error from method But we can ignore returned integer value by _
	//_, err := startWebServer(port)
	//fmt.Println(err)
}

// create custom Web server method with given port as parameter with Boolean type as return value
func startWebServer(port int) bool {
	fmt.Println("Starting server...")
	// do important thing
	fmt.Println("Server started on port", port)

	return true
}

// create custom Web server method with given port as parameter with Error type as return value
//func startWebServer(port int) error {
//	fmt.Println("Starting server...")
//	// do important thing
//	fmt.Println("Server started on port", port)
//
//	return nil
//}

// create custom Web server method with given port as parameter with Error type as return value
//func startWebServer(port int) error {
//	fmt.Println("Starting server...")
//	// do important thing
//	fmt.Println("Server started on port", port)
//
//	// don't forget to import "errors" module
//	return errors.New("Something want wrong")
//}


// create custom Web server method with given port as parameter with integer type and Error type as return values (multiple return values)
//func startWebServer(port int) (int, error) {
//	fmt.Println("Starting server...")
//	// do important thing
//	fmt.Println("Server started on port", port)
//
//	return port, nil
//}
