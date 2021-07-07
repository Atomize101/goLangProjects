package main

import "fmt"

func main() {
	//var operator string
	var number1, number2 int
	fmt.Println("Please enter a number")
	fmt.Scanln(&number1)
	fmt.Println("Please enter the next number")
	fmt.Scanln(&number2)
	fmt.Printf("%d + %d", number1, number2)
}
