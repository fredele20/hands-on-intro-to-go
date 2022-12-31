package main

import (
	"fmt"
	"runtime"
)

// if-else statement
// parseOddsEvens returns two slices, one with the odd numbers and one with the even numbers
func parseOddsEvens(ints []int) (odds []int, evens []int) {
	// use a for-range loop to iterate over the incoming slice
	for _, n := range ints {
		if n%2 == 0 {
			evens = append(evens, n)
		} else {
			odds = append(odds, n)
		}
	}

	// use the modulo operator to check if the number is odd or even and add it to the appropriate slice
	return
}

func main() {
	// invoke and print the results of parseOddEvens
	odds, evens := parseOddsEvens([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(odds)
	fmt.Println(evens)

	// Switch case statement
	// os := runtime.GOOS
	// if os == "linux" || os == "darwin" || os == "unix" {
	// 	fmt.Println("*nix variant")
	// } else if os == "windows" {
	// 	fmt.Println("windows")
	// } else {
	// 	fmt.Printf("%s.\n", os)
	// }

	switch os := runtime.GOOS; os {
	case "linux", "darwin", "unix":
		fmt.Println("*nix variant")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Printf("%s.\n", os)

	}

}
