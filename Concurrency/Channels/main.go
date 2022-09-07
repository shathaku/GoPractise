package main

import (
	"fmt"
	"time"
)

var words = []string{
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

func produceWord(c chan string) {
	for _, word := range words {
		// fmt.Println("sending a word")
		c <- word
		fmt.Println(word)
	}
	
	close(c)
}

func consumeWord(worker int, c chan string, done chan bool) {
	for w := range c {
		time.Sleep(1 * time.Second)
		fmt.Printf("Word %v is consumed by worker %v.\n", w, worker)
	}
	done <- true
}

func main() {

	c := make(chan string)
	done := make(chan bool)

	go produceWord(c)
	// go consumeWord(c, done)

	for i := 1; i <=3; i++ {
		go consumeWord(i, c, done)
	}
	
	fmt.Println(<- done)
	fmt.Println(<- done)
	fmt.Println(<- done)
}