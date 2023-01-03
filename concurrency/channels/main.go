package main

import (
	"fmt"
)

// sum - calculates and print the sum of numbers
func sum(nums []int, ch chan<- int) {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	ch <- sum
}

func main() {

	nums := []int{1, 2, 3, 4, 5}

	ch := make(chan int)

	go sum(nums, ch)
	result := <-ch
	fmt.Println("Result: ", result)

	// buffered channel
	ch2 := make(chan string, 2)

	ch2 <- "James"
	ch2 <- "Victor"

	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}
