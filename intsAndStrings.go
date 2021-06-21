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

	fmt.Print("Hello, ")
	fmt.Print("world, \n")
	fmt.Print("new line, \n")

	age := 35
	name := "Chris"

	// Println %_ = format specifier
	fmt.Println("hello")
	fmt.Println("goodbye")                                //Println makes the new line automatically
	fmt.Println("my age is", age, "and my name is", name) // output variables in printing

	// Printf (formatted strings)
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name) // puts quotes around a string variable
	fmt.Printf("age is of type %T \n", age)                    // Prints the type of the variable
	fmt.Printf("You scored %0.1f points \n", 225.55)           //Prints a float

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved str is:", str)

}

// https://golang.org/ref/spec#Numeric_types
