// This file has the tutorials of flow control statements.
package main

import (
	"fmt"
	"math"
)

// If can have an init condition just like for.
func pow(x, n, lim float64) float64 {
	// Variables declared inside an if short statement are
	// available inside the corresponding else statement.
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {

	// For loop example.
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("sum from for loop is %d\n", sum)

	// Init and post statements are optional.
	// Making for the while loop if required.
	for sum < 1000 {
		sum += sum
	}
	fmt.Printf("sum is now %d\n", sum)

	// Omit the condition and it becomes an infinite loop.
	// for {}

	// If is a standard construct.
	if sum > 1000 {
		fmt.Printf("sum is greater than 1000\n")
	}

	// If init condition.
	fmt.Printf("If with init condition 3^2 and 3^3 returns are %g, %g\n", pow(3, 2, 20), pow(3, 3, 20))
}
