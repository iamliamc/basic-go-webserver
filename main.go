package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/pluralsight/webservice/basicExamples"
	"github.com/pluralsight/webservice/controllers"
	"github.com/pluralsight/webservice/models"
)

func main() {
	// TODO important package methods are only exposed if they have leading capitals
	// This external visibility is controlled by capitalizing the first letter of the item declared.
	/// All declarations, such as Types, Variables, Constants, Functions, etc.,
	/// that start with a capital letter are visible outside the current package.
	showExamples := true
	basicExamples.BasicTypeExamples(showExamples)
	basicExamples.BasicFlowExamples(showExamples)
	basicExamples.BasicObjectExamples(showExamples)

	if showExamples {
		u := models.User{
			ID:        2,
			FirstName: "Liam",
			LastName:  "Considine",
		}

		fmt.Println(u)
	}
	if !showExamples {
		port := 3003
		_, err := startWebServer(port, 2)
		fmt.Println(err)
	} else {
		fmt.Println("Change showExamples bool to start the server")
	}
}

func startWebServer(port int, numberOfRetries int) (int, error) {
	controllers.RegisterControllers()
	fmt.Println("Starting server...")
	fmt.Println("Server started on port", port)

	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	fmt.Println("Number of retries", numberOfRetries)
	return port, errors.New("something went wrong")
}
