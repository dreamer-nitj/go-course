package main

import (
	"fmt"
	"math"
)

// Abser is the interface
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// in the following line, v is a Vertex ( not *Vertex )
	// and does not implement Abser.
	// a = v

	fmt.Println(a.Abs())
}

// MyFloat is the type
type MyFloat float64

// Abs is the method
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

// Vertex is the struct
type Vertex struct {
	X, Y float64
}

// Abs is the method
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
