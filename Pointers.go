package main

import (
	"fmt"
)

func updateName(x *string) string { // When you see the *, it means the type is expecting a pointer
	*x = "wedge"
	return *x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	// group A types > strings, ints, bools, floats, arrays, structs
	// Non-Pointer Wrapper Values
	name := "Chris"

	//updateName(name)

	//fmt.Println("memory address of name is: ", &name) // memory address of name

	m := &name
	fmt.Println("memory address:", m)
	fmt.Println("value at memory address: ", *m)

	fmt.Println(name)
	updateName(m)
	fmt.Println(name)

	//fmt.Println(name)

	// group B types > slices, maps, functions
	// Pointer Wrapper Values
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu)
}
