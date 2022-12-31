package main

import "fmt"

// define whatAmI which takes in an argument of any type and returns information about the underlying value's type
func whatAmI(val interface{}) string {
	switch val.(type) {
	case int:
		return fmt.Sprintf("%v is an int", val)
	case string:
		return fmt.Sprintf("%v is a string", val)
	default:
		return fmt.Sprintf("%v is not supported type", val)
	}
}

func main() {
	// invoke whatAmI functioin
	fmt.Println(whatAmI(1))
	fmt.Println(whatAmI("hello"))
	fmt.Println(whatAmI(true))
}