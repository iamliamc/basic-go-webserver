package main

import "fmt"

const (
	first  = 1
	second = 2

	firstIota  = iota + 6
	secondIota = 2 << iota
)

// Iota resets between constant blocks
const (
	block = iota
	evolveBlock
	evolveBlock2
)

const (
	// Iota 0 + 6 = 6
	constantExpression = iota + 6
	// Iota 1 + 6 = 7
	repeatLastConstantExpression
)

func main() {
	exploreConstantValues()
	explorePointerValues()
	exploreValueTypes()
}

func exploreConstantValues() {
	fmt.Println("exploreConstantValues")
	// implicit type
	const c = 3
	fmt.Println(c + 0.1415)
	fmt.Println(first, second)

	// One times two, five times
	bigShiftExample := 1 << 5

	// 32 divided by two, five times
	downShiftExample := bigShiftExample >> 5
	fmt.Println(bigShiftExample, downShiftExample)

	// Iota examples
	fmt.Println(firstIota, secondIota)
	fmt.Println(block, evolveBlock, evolveBlock2)
	fmt.Println(constantExpression, repeatLastConstantExpression)
}

func explorePointerValues() {
	fmt.Println("explorePointerValues")

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
	fmt.Println("exploreValueTypes")
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
