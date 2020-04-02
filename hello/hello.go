package main

import (
	"fmt"

	"example.com/user/hello/morestrings"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(morestrings.ReverseRunes("Hello Go!"))
}
