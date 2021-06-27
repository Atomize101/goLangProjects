package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills function
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"omlette": 9.99, "clam chowder": 4.99},
		tip:   0,
	}

	return b
}

// format the bill
// receiver function
// creates a copy
func (b bill) format() string {
	fs := "Bill Breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v) // the -25 before the v adds 25 characters to sort of center the line
		total += v
	}

	//total
	fmt.Println(b.name)
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}
