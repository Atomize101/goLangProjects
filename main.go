package main

import "fmt"

func add(n1 int, n2 int) int {
	return n1 + n2
}

func subtract(n1 int, n2 int) int {
	return n1 - n2
}

func main() {
	//var operator string
	total := 0
	total2 := 0
	var number1, number2 int
	fmt.Println("Please enter a number")
	fmt.Scanln(&number1)
	fmt.Println("Please enter the next number")
	fmt.Scanln(&number2)
	total = add(number1, number2)
	total2 = subtract(number1, number2)
	fmt.Printf("%d + %d = %d \n", number1, number2, total)
	fmt.Printf("%d - %d = %d \n", number1, number2, total2)
}
