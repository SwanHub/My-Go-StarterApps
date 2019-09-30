package main

import (
	"fmt"
	"strconv"
)

// define person struct ==> similar to CLASSES
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
} // use "prettier" in vs code

// the two types of methods with structs
// Greeting method (value receiver) this is the getter method
func (p Person) greet() string { // p is an identifier, could be anything === p is instead of this / self
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) // can't have mismatched
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver) this is the setter method
func (p *Person) getMarried(spouseLastName string) { // if you want to mutate with a class function, use '*'
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using a struct
	// person1 := Person{firstName: "Samantha", lastName: "Smithereens", city: "SF", gender: "m", age: 3}

	// Alternative
	person1 := Person{"Sam", "Smithie", "SF", "f", 3}
	person2 := Person{"Bob", "Johson", "SF", "m", 3}

	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)

	person1.hasBirthday()
	person1.getMarried("Williams")
	person2.getMarried("Barry")

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())

}
