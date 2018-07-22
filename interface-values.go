package main

import (
	"fmt"
	"math"
)

// I is the interface
type I interface {
	M()
}

// T is the interface
type T struct {
	S string
}

// M is the method
func (t *T) M() {
	fmt.Println(t.S)
}

// F is the type
type F float64

// M is the method
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
