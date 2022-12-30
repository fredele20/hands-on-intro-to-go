package main

import "fmt"


func main() {
	// declare a slice string with the make function
	names := make([]string, 0)

	// append 3 names to the slice
	names = append(names, "John")
	names = append(names, "Jane")
	names = append(names, "Mary")

	// print the slice
	fmt.Println(names)

	// slice the slice using syntax alice(low:high)
	fmt.Println(names[1:3])// [Jane Mary]
	fmt.Println(names[1:])// [Jane Mary]
	fmt.Println(names[:1])// [John]
	fmt.Println(names[:])// [John Jane Mary]
}