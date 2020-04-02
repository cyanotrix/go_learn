// The main package of my tutorial.
package main

import (
	"example.com/user/hello/morestrings"

	"fmt"

	"github.com/google/go-cmp/cmp"
)

// Same type parameters can be clubbed.
func add(a, b int) int {
	return a + b
}

// Function can have multiple returns.
func swap(a, b string) (string, string) {
	return b, a
}

// Naked return values example.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// Hello world.
	fmt.Println("hello world")

	// Packages and reverserunes function.
	fmt.Println(morestrings.ReverseRunes("Hello Go!"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))

	// Sum and add.
	fmt.Printf("Sum of 45 and 29 is %d\n", Sum(45, 29))
	fmt.Printf("Add of 786 and 233 is %d\n", add(786, 233))

	// Swap.
	a, b := swap("hello", "world")
	fmt.Printf("Swapping 2 string hello and world %q, %q\n", a, b)

	// Split.
	x, y := split(43)
	fmt.Printf("Split of integer 43 are %d, %d\n", x, y)
}
