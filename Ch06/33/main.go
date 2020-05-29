package main

import (
	"github.com/pluralsight/webservice/controllers"
	"net/http"
)

func main() {
	// register front controllers
	controllers.RegisterControllers()

	// initialize Http server listen on 3000 port and use default ServeMax way which handler by front.go file
	http.ListenAndServe(":3000", nil)
}
