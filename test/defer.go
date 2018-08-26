package main

import (
	"fmt"

	"github.com/dreamer-nitj/stringutils"
	// "github.com/golang/example/stringutil"
)

func main() {
	defer fmt.Println("World")

	fmt.Println("Hello ")
	fmt.Printf(stringutils.Reverse("!oG ,olleH"))
}
