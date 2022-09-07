package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))

	for i, w := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, w), &wg)
	}

	wg.Add(1)
	printSomething("this is the second line", &wg)
	wg.Wait()
}

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}