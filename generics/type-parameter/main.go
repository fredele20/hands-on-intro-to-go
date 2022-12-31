package main

import "fmt"

func sumInts(a, b int) int {
	return a + b
}

func sumFloats(a, b float64) float64 {
	return a + b
}

// create generic sum functioin with type parameter T constrained to int and float64 types
func sum[T ~int | ~float64](a, b T) T {
	return a + b
}

type specialInt int

func main() {
	// non-generic sum int function
	// fmt.Println(sumInts(1, 2))

	// non-generic sum float function
	// fmt.Println(sumFloats(1.3, 2.2))

	// call on generic sum function
	// fmt.Println(sum[int](1, 2))
	// fmt.Println(sum[float64](1.3, 2.2))

	// define a compatible custom type call on generic sum function with it
	one := specialInt(1)
	two := specialInt(2)
	fmt.Println(sum(one, two))
}
// list is a singly-linked list that holds values of any type
type list[T any] struct {
	next *list[T]
	val T
}
