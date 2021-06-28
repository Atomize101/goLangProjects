package main

import "fmt"

func main() {
	mybill := newBill("Chris' Bill")

	mybill.addItem("soup", 4.50)
	mybill.addItem("apple", 1.99)
	mybill.addItem("clam chowder", 11.99)
	mybill.addItem("ceasar salad", 9.99)
	mybill.updateTip(10)

	fmt.Println(mybill.format())
}
