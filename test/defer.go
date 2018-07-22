package main

import (
	"fmt"

	"github.com/dreamer-nitj/stringutils"
)

func main() {
	defer fmt.Println("World")

	fmt.Println("Hello ")
	fmt.Printf(stringutils.Reverse("!oG ,olleH"))
}
