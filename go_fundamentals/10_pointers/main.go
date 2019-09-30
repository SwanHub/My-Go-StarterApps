package main

import "fmt"

func main() {
	a := 5
	b := &a //assigns b to a pointer of a ==> the memory addess is returned.

	fmt.Println(a, b)
	fmt.Printf("%T\n", a) // int
	fmt.Printf("%T\n", b) // *int (different values)

	// Use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)  // points to the value
	fmt.Println(b)  // points to the memory address of a aka &a
	fmt.Println(*b) // points to the value of a
	fmt.Println(&a) // memory address

	// you might have to pass the address instead of the data itself...
}
