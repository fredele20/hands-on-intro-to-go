package main

import "fmt"

// create a numeric interface with a type set
type numeric interface {
	~int | ~float64
	grow()
}

// create generic sum functioin with type parameter T constrained to int and float64 types
func sum[T numeric](a, b T) T {
	return a + b
}

type specialInt int
func (si specialInt) grow() {}

func main() {
	one := specialInt(1)
	two := specialInt(2)
	fmt.Println(sum(one, two))
}
