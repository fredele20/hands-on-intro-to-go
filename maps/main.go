package main

import "fmt"

type author struct {
	name string
}


func main() {
	// initialize the map with make
	authors := map[string]author{
		"tm": {name: "Toni Morrison"},
		"ma": {name: "Marcus Aurelius"},
	}

	// print the map with fmt.Printf
	// fmt.Printf("%#v\n", authors)

	// read a value from the map with a known key
	fmt.Println(authors["tm"])

	// read the value for a non-existent key
	fmt.Println("JR: ", authors["jr"])

	// check when a key is present in the map
	a, ok := authors["jr]"]
	fmt.Printf("a = %v, ok = %v\n", a, ok)

	// updating a value in maps
	// change the value of an author in the map
	authors["tm"] = author{name: "James Logan"}
	fmt.Printf("%v\n", authors)

	// delete a key from map
	delete(authors, "tm")
	fmt.Printf("%v\n", authors)

}