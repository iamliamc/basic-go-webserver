package main

import (
	"errors"
	"fmt"

	basicexamples "github.com/pluralsight/webservice/basicExamples"
	"github.com/pluralsight/webservice/controllers"
	"github.com/pluralsight/webservice/models"
)

func main() {
	// TODO important package methods are only exposed if they have leading capitals
	// This external visibility is controlled by capitalizing the first letter of the item declared.
	/// All declarations, such as Types, Variables, Constants, Functions, etc.,
	/// that start with a capital letter are visible outside the current package.

	basicexamples.Examples(false)
	u := models.User{
		ID:        2,
		FirstName: "Liam",
		LastName:  "Considine",
	}

	fmt.Println(u)
	port := 1000
	_, err := startWebServer(port, 2)
	fmt.Println(err)
}

func startWebServer(port int, numberOfRetries int) (int, error) {
	fmt.Println("Starting server...")
	controllers.RegisterControllers()
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
	return port, errors.New("Something went wrong")
}
