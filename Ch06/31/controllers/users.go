package controllers

import (
	"net/http"
	"regexp"
)

// Create User controller with Struct type (keep it empty)
type UserController struct {
	userIDPattern *regexp.Regexp
}

// Now bind method to created controller, after func keyword will set type which wanted to bind this method
// on C# there extensions method principle, Example => void Test(this UserController, port int)
func (uc UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// write string message to response with bytes slice
	w.Write([]byte("Hello from the User Controller!"))
}

// use Go practice to make Constructor method for custom class type to create UserController constructor
// will return pointer of create new User Controller instance
func newUserController() *UserController {
	return &UserController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}