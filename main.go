package main

import (
	"fmt"
	"sync"
)

func goRoutine(s string,  wg *sync.WaitGroup) {
	defer wg.Done() // decrement wg by one after this function completes
	println(s)
}

func main() {
	var wg sync.WaitGroup
	words := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"epsilon",
		"zeta",
		"eta",
		"theta",
		"iota",
		"kappa",
		"lambda",
		"mu",
		"nu",
	}


	wg.Add(len(words))

	for i, s := range words {
		goRoutine(fmt.Sprintf("%d: %s", i, s), &wg)
	}

	wg.Wait()

	wg.Add(1)
	goRoutine("Second thing to be printed!", &wg)
}