package main

import "fmt"

func main() {
	// arrays in go have a fixed length, cannot change
	// var ages [3]int = [3]int{20, 25, 30} // array structure
	var ages2 = [3]int{20, 25, 30} // this is the same as the array above

	names := [4]string{"Chris", "Sarah", "Demario", "Mike"}
	names[1] = "Bob"

	fmt.Println(ages2, len(ages2)) // prints out the array and the length
	fmt.Println(names, len(names))

	// slices (use arrays under the hoood)
	var scores = []int{100, 50, 60} // when you dont put in a number in [] it creates a slice
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]   // leaving out the second value stores the first up until the end
	rangeThree := names[:3] // starts from the beginning of slice and stores up until value

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Silly")
	fmt.Println(rangeOne)
}
