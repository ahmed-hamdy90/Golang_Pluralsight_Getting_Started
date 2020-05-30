package main

// create HTTPRequest as struct type
type HTTPRequest struct {
	Method string
}

func main() {
	// create instance from HTTPRequest
	r := HTTPRequest {Method: "GET"}

	// In Go we don't need to add break after every case block as switch under Go is implicit switch
	switch r.Method {
	case "GET":
		// User built-in println method
		println("Get request")
	case "DELETE":
		println("Delete request")
	case "POST":
		println("Post request")
	case "PUT":
		println("Put request")
	}

	// You can still use break after every case block => will work too
	switch r.Method {
	case "GET":
		println("Get request")
		break
	case "DELETE":
		println("Delete request")
		break
	case "POST":
		println("Post request")
		break
	case "PUT":
		println("Put request")
		break
	}

	// If you need to full two cases blocks, use fallthrough keyword after first case block
	switch r.Method {
	case "GET":
		println("Get request")
		fallthrough
	case "DELETE":
		println("Delete request")
	case "POST":
		println("Post request")
	case "PUT":
		println("Put request")
	}

	//r := HTTPRequest {Method: "HEAD"}

	// If you need to handle false case, use default block
	//switch r.Method {
	//case "GET":
	//	println("Get request")
	//case "DELETE":
	//	println("Delete request")
	//case "POST":
	//	println("Post request")
	//case "PUT":
	//	println("Put request")
	//default:
	//	println("Unhandled method")
	//}
}