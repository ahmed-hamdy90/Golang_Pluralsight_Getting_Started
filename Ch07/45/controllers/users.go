package controllers

import (
	"encoding/json"
	"github.com/pluralsight/webservice/models"
	"net/http"
	"regexp"
	"strconv"
)

// Create User controller with Struct type (keep it empty)
type UserController struct {
	userIDPattern *regexp.Regexp
}

// Now bind method to created controller, after func keyword will set type which wanted to bind this method
// on C# there extensions method principle, Example => void Test(this UserController, port int)
func (uc UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// according to Request's path and verb determine which method will be called
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		// Else condition will be /users/(id+) URL so must make sure is already same URL
		// also getting ID value from URL => using detect regex match groups for (id+)
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}

		// convert given id value from Request's URL to integer value
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}

		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		case http.MethodPut:
			uc.put(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

// Create method to Handle request for Get all Users
func (uc UserController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), w)
}

// Create method to Handle request for Get User by id value
func (uc UserController) get(id int, w http.ResponseWriter) {
	// get user by given id
	u, err := models.GetUserByID(id)
	// check if GetUserByID method return error or not
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encodeResponseAsJSON(u, w)
}

// Create method to Handle request for add new user
func (uc UserController) post(w http.ResponseWriter, r *http.Request) {
	// Parse given POST body
	u, err := uc.parseRequest(r)
	// check if parseRequest method return error or not
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}

	// add parsed user
	// Notice: use = equal sign rather than := implicit type sign as both u and err was not new variables(already created from above call method)
	u, err = models.AddUser(u)
	// check if AddUser method return error or not
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	encodeResponseAsJSON(u, w)
}

// Create method to Handle request for edit exists user
func (uc UserController) put(id int, w http.ResponseWriter, r *http.Request) {
	// Parse given PUT body
	u, err := uc.parseRequest(r)
	// check if parseRequest method return error or not
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}

	// make sure given parsed user has same ID value of given id value
	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}

	// update parsed user
	// Notice: use = equal sign rather than := implicit type sign as both u and err was not new variables(already created from above call method)
	u, err = models.UpdateUser(u)
	// check if UpdateUser method return error or not
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	encodeResponseAsJSON(u, w)
}

// Create method to Handle request for remove exists user
func (uc UserController) delete(id int, w http.ResponseWriter) {
	// remove user by given id
	err := models.RemoveUserById(id)
	// check if RemoveUserById method return error or not
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Create method to parse Body data into both POST/PUT request
func (uc UserController) parseRequest(r *http.Request) (models.User, error) {
	// decode Request's body using Decode method under Json module  => first get instance from Decoder class
	dec := json.NewDecoder(r.Body)

	// create variable from User class to decode body data on it
	var u models.User
	err := dec.Decode(&u)
	// check if Decode method return error or not
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

// use Go practice to make Constructor method for custom class type to create UserController constructor
// will return pointer of create new User Controller instance
func newUserController() *UserController {
	return &UserController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}