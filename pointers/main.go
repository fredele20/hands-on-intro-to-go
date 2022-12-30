package main

import "fmt"


func main() {
	var a *int
	
	b := 100

	a = &b

	fmt.Println(a)
	
	fmt.Println(*a)
}