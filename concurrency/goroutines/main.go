package main

import (
	"fmt"
	"time"
)

// sum - calculates and print the sum of numbers
func sum(nums []int) {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println("Result: ", sum)
}

func main() {

	nums := []int{1, 2, 3, 4, 5}
	go sum(nums)
	time.Sleep(100 * time.Millisecond)
	
}