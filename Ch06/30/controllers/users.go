package controllers

import "net/http"

// Create User controller with Struct type (keep it empty)
type UserController struct {}

// Now bind method to created controller, after func keyword will set type which wanted to bind this method
// on C# there extensions method principle, Example => void Test(this UserController, port int)
func (uc UserController) ServeHttp(w http.ResponseWriter, r *http.Request) {
	// write string message to response with bytes slice
	w.Write([]byte("Hello from the User Controller!"))
}