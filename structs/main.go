package main

import "fmt"

// define a struct type of person
type person struct {
	first string
	last string
}

// function structs
// fullname returns the full name of the person
func (p person) fullName() string {
	return p.first + " " + p.last
}

//changeName changes the first and last name of the person
func (p *person) changeName(first, last string) {
	p.first = first
	p.last = last
}

// struct embedding
// define author and embed person
type author struct {
	person
	penName string
}

// override fullName method for author
func (a author) fullName() string {
	return fmt.Sprintf("%s (%s)", a.person.fullName(), a.penName)
}

func main() {

	// initialize person
	p := person{
		first: "Victor",
		last: "Ishola",
	}

	// fmt.Printf("%#v\n", p)
	fmt.Println(p.fullName())

	p.changeName("Fredel", "Olad")
	fmt.Println(p.fullName())

	// initialize and print author's full name
	a := author{
		person: person{
			first: "James",
			last: "Logan",
		},
		penName: "Smart",
	}

	fmt.Println(a.fullName())
	
}