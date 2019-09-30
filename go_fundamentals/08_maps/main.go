package main

import "fmt"

// MAPS ARE OBJECTS
func main() {
	// Define map (key value pair(?)) defining an object
	// emails := make(map[string]string) // defines a map, key and value both string

	// Assign key
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"

	// declare map and add key values
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	emails["Mike"] = "mike@gmail.com" // auto sorts by alphabetical => key...

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"]) // use a better key example.

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
