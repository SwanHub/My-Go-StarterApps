package main

import "fmt"

// range == loop through arrays, maps, slices
func main() {

	ids := []int{33, 45, 21, 33, 4}

	// loop through ids using range
	for i, id := range ids { // gets the index and the value
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index ==> underscore makes up for unused index values.
	for _, id := range ids { //
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}

	for _, v := range emails {
		fmt.Println("Emails: " + v)
	}
}
