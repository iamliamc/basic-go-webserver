package basicexamples

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

func Examples(showOutput bool) {
	if showOutput {
		exploreStructs()
		exploreMaps()
		exploreArrays()
		exploreSlices()
		exploreConstantValues()
		explorePointerValues()
		exploreValueTypes()
	}
}

func exploreStructs() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.ID = 1
	u.FirstName = "Liam"
	u.LastName = "Considine"
	fmt.Println(u.FirstName)

	u2 := user{ID: 2, FirstName: "Ciaran", LastName: "Considine"}
	fmt.Println(u2)
}

func exploreMaps() {
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 27
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)
}

func exploreSlices() {
	arr := [3]int{1, 2, 3}

	// Create a slice of the array from the beginning of that collection to the end
	slice := arr[:]

	arr[1] = 42
	slice[2] = 27

	// Shows that the slice has a pointer to the array
	fmt.Println(arr, slice)

	slice2 := []int{1, 2, 3}
	slice2 = append(slice2, 4, 42, 27)
	fmt.Println(slice2)

	s3 := slice2[1:]
	s4 := slice2[:2]
	s5 := slice[1:2]
	fmt.Println(s3, s4, s5)
}

func exploreArrays() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
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
