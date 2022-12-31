package main

import "fmt"


func main() {
	// Simple for loop
	// declare a string to iterate over
	// s := "Hello"

	// iterate over the string with basic for loop
	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(i, ":", string(s[i]))
	// }

	// checking the continuity of a vairiable
	// initialize a variable to check if the loop should continue
	// n := 5

	// iterate while the condition is true
	// for n > 0 {
	// 	fmt.Println(n)
	// 	n--
	// }

	// Using loop with the range keyword
	// initialiaze a slice of ints
	nums:= []int{100, 200, 300}


	// use for-range to iterate over the slice and print each value
	for i, n := range nums {
		fmt.Println(i, n)
	}

	// declare a map of string to ints
	m := map[string]int{"one": 1, "two": 2, "three":3}

	// use for-range to iterate over the map and print each key/value pair
	for k, v := range m {
		fmt.Println(k, v)
	}

	// infinite loop
	// for {

	// }
}