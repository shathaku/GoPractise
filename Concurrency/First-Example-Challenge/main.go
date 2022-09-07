package main

import (
	"fmt"
	"sync"
)

var msg string

func main() {
	msg = "Hello, world"

	var wg sync.WaitGroup

	wg.Add(1)
	go updateMessage("Hello, universe", &wg)
	// printMessage()
	wg.Wait()

	wg.Add(1)
	go updateMessage("Hello, cosmos", &wg)
	// printMessage()
	wg.Wait()

	wg.Add(1)
	go updateMessage("Hello, world", &wg)
	// printMessage()
	wg.Wait()
}

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

// func printMessage() {
// 	fmt.Println(msg)
// }