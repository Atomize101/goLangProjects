package main

import (
	"fmt"
)

var score = 99.5 // If this was put in the main function it would not work
// It has to be declared in the root

func main() {

	sayHello("Chris")

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()
}
