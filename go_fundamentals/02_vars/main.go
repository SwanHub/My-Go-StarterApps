package main

import "fmt"

// var name = "Brad" // global

func main() {
	// you HAVE to use a variable if it's declared.
	// Main types
	// string, bool, int, int8, int16, int-64 (length), uint, uint8-64, byte (alias for uint8), rune (alias for int32)
	// float32, 64
	// complex64, 128

	// var
	// var name = "Brad"  // string implicitly
	var age int32 = 37 //int implicity ==> you can declare specific int though...
	const isCool = true
	// isCool = false can't do this, because const can't be redefined.

	// shorthand
	// name := "Brad"
	size := 1.3 // float64 is implicit (var size float32 = 2.3)
	// email := "brad@gmail.com"

	name, email := "Brad", "brad@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}
