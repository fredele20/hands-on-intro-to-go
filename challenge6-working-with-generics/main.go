package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// part 1: print function refactoring
func printAny[T any](a T) { fmt.Println(a) }

// non-generic print functions

func printString(s string) { fmt.Println(s) }

func printInt(i int) { fmt.Println(i) }

func printBool(b bool) { fmt.Println(b) }

// part 2: sum function refactoring

type numeric interface {
	constraints.Integer | constraints.Float
}

func sumRefactor[T numeric](nums ...T) T {
	var sum T
	for _, n := range nums {
		sum += n
	}
	return sum
}

// sum sums a slice of any type
func sum(numbers []interface{}) interface{} {
	var result float64
	for _, n := range numbers {
		switch n.(type) {
		case int:
			result += float64(n.(int))
		case float32, float64:
			result += n.(float64)
		default:
			continue
		}
	}
	return result
}

func main() {
	printAny("Hello")
	printAny(2)
	printAny(true)

	fmt.Println(sumRefactor(2, 6, 7, 2))

}
