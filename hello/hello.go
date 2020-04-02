package main

import (
	"example.com/user/hello/morestrings"

	"fmt"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(morestrings.ReverseRunes("Hello Go!"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
