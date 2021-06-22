package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))     // returns new string all uppercase
	fmt.Println(strings.Index(greeting, "ll")) // gets the position of the second argument
	fmt.Println(strings.Split(greeting, " "))  // Splits the string based on second argument

	// the original value is unchanged
	fmt.Println("original string value =", greeting)

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages) // Will slice integers and sort them
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"Bob", "Chris", "Sarah", "Demario", "Jim"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Bob")) // Gives us the position in the slice

}
