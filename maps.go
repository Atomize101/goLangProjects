package main

import (
	"fmt"
)

// maps allow us to store key and value pairs

func main() {
	menu := map[string]float64{ //key will be a string, value will be a float
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{ // map keys will be ints, values will be strings
		502897: "Chris",
		889254: "Sarah",
		425365: "Bob",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[502897])

	phonebook[889254] = "Demario"
	fmt.Println(phonebook)

	phonebook[425365] = "Jim"
	fmt.Println(phonebook)

}
