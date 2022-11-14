package basicExamples

func BasicFlowExamples(showOutput bool) {
	if showOutput {
		basicLoopWithBreak()
		basicLoopWithContinue()
		basicLoopWithPostClause()
		basicInfiniteLoop()
		basicLoopOverCollections()
	}
}

func basicLoopWithBreak() {
	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			break
		}
	}
}

func basicLoopWithContinue() {
	var i int
	for i < 5 {
		println(i)
		i++
		if i == 3 {
			continue
		}
	}
}

func basicLoopWithPostClause() {
	for i := 0; i < 5; i++ {
		println(i)
	}
}

func basicInfiniteLoop() {
	var i int
	for {
		if i == 5 {
			break
		}
		println(i)
		i++
	}
}

func basicLoopOverCollections() {
	slice := []int{1, 2, 3}
	// ugly version with indexer
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	for i, v := range slice {
		println(i, v)
	}

	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for k, v := range wellKnownPorts {
		println(k, v)
	}
}

// func panicConcept() {
// 	// And how to recover from it
// 	panic("Something bad just happened")
// }
