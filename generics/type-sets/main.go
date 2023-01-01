package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// create a numeric interface with a type set
type numeric interface {
	// ~int | ~float64
	// implementing constraints
	constraints.Integer | constraints.Float
	grow()
}

// create generic sum functioin with type parameter T constrained to int and float64 types
func sum[T numeric](a, b T) T {
	return a + b
}

type specialInt int

func (si specialInt) grow() {}

// equal returns true if a and b are equal
func equal[T comparable](a, b T) bool {
	return a == b
}

func main() {

	// invoke equal with comparable types
	// fmt.Println("equal (1, 1):", equal(1, 1))
	// fmt.Println("equal(\"one\", \"two\"):", equal("one", "two"))

	// invoke with a custom type
	type c struct{ f string }
	fmt.Println("equal(c{f: \"a\"}, c{f: \"a\"}):", equal(c{f: "a"}, c{f: "a"}))
	fmt.Println("equal(c{f: \"a\"}, c{f: \"b\"}):", equal(c{f: "a"}, c{f: "b"}))

	// one := specialInt(1)
	// two := specialInt(2)
	// fmt.Println(sum(one, two))
}
