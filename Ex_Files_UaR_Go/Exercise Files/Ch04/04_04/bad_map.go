package main

import (
	"fmt"
)

func main() {
	// this is an error as the memory is not allocated
	var m map[string]int
	// this below way is the correct way to declare a map
	// m := make(map[string]int)
	m["key"] = 42
	fmt.Println(m)
}
