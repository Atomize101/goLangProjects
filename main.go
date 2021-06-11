package main

import "fmt"

func main() {
	nameOne := "Chris"           // short hand way of assigning a variable
	var nameTwo string = "Sarah" // More traditional method
	fmt.Println(nameOne, nameTwo)

	var ageOne int = 20
	ageTwo := 40 // Same method for ints
	fmt.Println(ageOne, ageTwo)

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 99999999999999999999
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)

}

// https://golang.org/ref/spec#Numeric_types
