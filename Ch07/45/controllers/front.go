package controllers

import (
	"encoding/json"
	"net/http"
)

// Create method responsible for initialize possible controllers plus possible routes which this controllers handler
func RegisterControllers() {
	// register User Controller using implicit type
	uc := newUserController()

	// add possible routes which User Controller handler
	// As UserController implement ServeHttp method which abstract method under http.Handler interface
	// So UserController will act as Handler instance with any implement keyword into UserController struct type declaration
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

// Create method responsible for encoding any data as JSON for HTTP response
func encodeResponseAsJSON(data interface{}, w http.ResponseWriter) {
	// use Encode method under Json module to encode response => first get instance from Encoder class
	enc := json.NewEncoder(w)
	enc.Encode(data)
}