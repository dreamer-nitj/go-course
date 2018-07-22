package main

import (
	"fmt"
	"math"
)

// MyFloat is the type
type MyFloat float64

// Abs is the method
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
