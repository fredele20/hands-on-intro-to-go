package main

import "fmt"

func cleanup(msg string) {
	fmt.Println(msg)
}

func main() {
	// defer function calls
	defer cleanup("first")
	defer cleanup("second")

	fmt.Println("Working with main....")

	// defer recovery
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from panic", err)
		}
	}()

	// panic
	panic("Something bad happened")
}