package main

import (
	"errors"
	"fmt"

	"github.com/pluralsight/webservice/models"
)

func main() {
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
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
	return port, errors.New("Something went wrong")
}
