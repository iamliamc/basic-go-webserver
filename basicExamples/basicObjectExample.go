package basicExamples

type Person struct {
	ID        int
	FirstName string
	LastName  string
}

func (p *Person) addMiddleName(middleName string) {
	println(p.FirstName, middleName, p.LastName)
}

func callMethodOnObject() {
	person := Person{ID: 1, FirstName: "Liam", LastName: "Considine"}
	person.addMiddleName("Joseph")
}

func BasicObjectExamples(showExamples bool) {
	if showExamples {
		callMethodOnObject()
	}
}
