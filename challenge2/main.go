package main

import "fmt"

var val = "Victor"

func printGlobalVal() {
	fmt.Println(val)
	fmt.Printf("%v=%T\n", val, val)
}

func updateGlobalVal() {
	val = "Fredel"
	fmt.Println(val)
}

func main() {

	val := 5
	fmt.Printf("%v=%T\n", val, val)
	fmt.Println(val)

	*(&val) = 100
	fmt.Println("Global val = ", val)

	
	printGlobalVal()
	updateGlobalVal()

	fmt.Printf("%T, local &val = %v\n", &val, val)
}