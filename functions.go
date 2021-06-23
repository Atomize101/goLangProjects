package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good Morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleAREA(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreeting("Chris")
	sayGreeting("Bob")
	sayBye("Sarah")

	cycleNames([]string{"Jim", "Brenda", "Mike"}, sayGreeting)
	cycleNames([]string{"Jim", "Brenda", "Mike"}, sayBye)

	a1 := circleAREA(10.5)
	a2 := circleAREA(15)

	fmt.Println(a1, a2)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)
}
