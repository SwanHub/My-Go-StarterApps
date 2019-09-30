package main

import (
	"fmt"
	"math"

	"github.com/SwanHub/go_tests/03_packages/strutil"
) //use prettier to automatically save imports this style...

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("olleh"))
	// fmt.Println(math.Ceil(2.7))
}
