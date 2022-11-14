package basicExamples

func BasicFlowExamples(showOutput bool) {
	if showOutput {
		basicLoop()
	}
}

func basicLoop() {
	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			break
		}
	}
}
