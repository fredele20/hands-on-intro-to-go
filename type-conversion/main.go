package main

import "fmt"

func main() {
	a := 42.5
	b := uint(a)
	fmt.Printf("%v=%T, %v=%T\n", a, a, b, b)
}