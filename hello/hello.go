// The main package of my tutorial.
package main

import (
	"math/cmplx"

	"example.com/user/hello/morestrings"

	"fmt"

	"github.com/google/go-cmp/cmp"

	"math"
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

// Variables with initializers
var i, j int = 1, 2

// Basic types example.
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// Constant example.
const Pi = 3.14159265

func main() {
	// Hello world.
	fmt.Println("hello world")

	// Packages and reverserunes function.
	fmt.Println(morestrings.ReverseRunes("Hello Go!"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))

	// Sum and add.
	fmt.Printf("Sum of 45 and 29 is %d\n", morestrings.Sum(45, 29))
	fmt.Printf("Add of 786 and 233 is %d\n", add(786, 233))

	// Swap.
	a, b := swap("hello", "world")
	fmt.Printf("Swapping 2 string hello and world %q, %q\n", a, b)

	// Split.
	x, y := split(43)
	fmt.Printf("Split of integer 43 are %d, %d\n", x, y)

	// Var with inits.
	fmt.Printf("i, j = %d, %d\n", i, j)

	// Short variable declaration, type is implicit.
	k := true
	fmt.Printf("short variable declaration k is %t", k)

	// Basic datatypes.
	fmt.Printf("ToBe Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("MaxInt Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("z Type: %T, value: %v\n", z, z)

	// Type conversions.
	x, y = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	g := uint(f)
	fmt.Printf("x, y and g are %v, %v, %v\n", x, y, g)

	// Constant example.
	const World = "世界"
	fmt.Printf("Hello %s\n", World)
	fmt.Printf("Happy %.8f day\n", Pi)
}
