// The main package of my tutorial.
package main

import (
	"example.com/user/hello/morestrings"

	"fmt"

	"github.com/google/go-cmp/cmp"
)

func add(a, b int) int {
	return a + b
}

func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	fmt.Println("hello world")
	fmt.Println(morestrings.ReverseRunes("Hello Go!"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))

	fmt.Printf("Sum of 45 and 29 is %d", Sum(45, 29))
	fmt.Printf("\n")
	fmt.Printf("Add of 786 and 233 is %d", add(786, 233))
	fmt.Printf("\n")

	a, b := swap("hello", "world")
	fmt.Printf("Swapping 2 string hello and world %q, %q", a, b)
}
