package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	defer fmt.Println("World")

	fmt.Println("Hello ")
	rand.Seed(time.Now().UnixNano())
	fmt.Println("The random number is", rand.Intn(100))
}
