package main

import "fmt"

func add(n1 int, n2 int) int {
	return n1 + n2
}

func subtract(n1 int, n2 int) int {
	return n1 - n2
}

var result int

func main() {
	var operator string
	var number1, number2 int

	fmt.Println("Please enter a number")
	fmt.Scanln(&number1)
	fmt.Println("Please enter the next number")
	fmt.Scanln(&number2)
	fmt.Print("Please enter Operator (+,-,/,%,*):")
	fmt.Scanln(&operator)
	fmt.Printf("%d %s %d = %d \n", number1, operator, number2, result)

	switch operator {
	case "+":
		result = add(number1, number2)
		fmt.Println(result)
	}
}

// she beat me 7/9/2021
