package main

import (
	"fmt"
)

func main() {
	// variable declaration: explicit
	var anInteger int = 42
	var anInt int
	anInt = 33

	fmt.Println(anInteger)
	fmt.Println(anInt)

	// variable declaration: implicit
	anInteger2 := 44
	fmt.Println(anInteger2)

	// constant declaration explicit
	const anIntConst int = 23
	fmt.Println(anIntConst)

	// constant declaration implicit
	const anIntConst2 = 22
	fmt.Println(anIntConst2)
}
