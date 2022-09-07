package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

func main() {
	c1 := circle{ 5.2 }
	s1 := square{ 4.8 }

	shape := []shape{c1, s1}

	fmt.Println(shape[0].area())
	fmt.Println(shape[1].area())
}