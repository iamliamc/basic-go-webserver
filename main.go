package main

import "fmt"

func main() {
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
