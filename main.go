package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// func greet() string {
// 	return "Hello"
// }

// func greetWithName(name string) string {
// 	return "Hello, " + name + "!"
// }

func greetWithNameAndAge(name string, age int) (greeting string) {
	greeting = "Hello, my name is " + name + " and I am " + strconv.Itoa(age) + " Years old."
	return
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by 0")
	}

	return (a / b), nil
}

func main() {
	fmt.Println("Hello Gopher!")

	// fmt.Println(greetWithName("Victor"))
	fmt.Println(greetWithNameAndAge("Victor", 23))

	fmt.Println(divide(5, 0))
	fmt.Println(divide(9, 3))

	fmt.Printf("Today is %s\n", time.Now().Weekday())
}
