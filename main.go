package main

import "fmt"

func main() {
	x := 0
	for x < 5 {
		fmt.Println("value of x:", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of i:", i)
	}

	names := []string{"Bob", "Chris", "Sarah", "Demario", "Jim"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names { //This is like for in
		fmt.Printf("the value at index %v is %v \n", index, value)
	}

	for _, value := range names { //This is like for in
		fmt.Printf("the value is %v \n", value)
		value = "new string"
	}

	fmt.Println(names)

}
