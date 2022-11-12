package main

import "fmt"

func main() {
	// pointer operator
	var firstName *string = new(string)

	// dereference pointer operator
	*firstName = "Arthur"

	fmt.Println(*firstName)

	lastName := "Considine"
	ptr := &lastName
	fmt.Println(ptr, *ptr)

	// pointer is the same value is different
	lastName = "Hazle"
	fmt.Println(ptr, *ptr)
}

func exploreValueTypes() {
	fmt.Println("Hello from a module, Gophers!")
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.141516
	fmt.Println(f)

	firstName := "Liam"
	fmt.Println(firstName)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)
}
