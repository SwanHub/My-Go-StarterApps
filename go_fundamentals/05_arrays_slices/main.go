package main

import "fmt"

// Arrays need to be a fixed length (slices deal with arrays without a fixed type)
// they dont' use commas
func main() {
	// Arrays
	// var fruitArr [2]string

	// Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// declare and assign ==> shorthand
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	// Slices => not declaring the length of the array.
	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2]) //non inclusive, [start:stop].
}
