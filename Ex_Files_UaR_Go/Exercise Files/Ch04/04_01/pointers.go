package main

import "fmt"

func main() {
	// declaring a pointer that can point to any int
	var p *int

	// if we try to print the value here, we will get nil pointer dereference runtime error
	// fmt.Println("Value of p:", *p)

	if p != nil {
		fmt.Println("Value of p:", *p)
	} else {
		fmt.Println("p is nil")
	}

	var v int = 42
	p = &v
	if p != nil {
		fmt.Println("Value of p:", *p)
	} else {
		fmt.Println("p is nil")
	}

	var value1 float64 = 42.13
	pointer1 := &value1
	fmt.Println("Value 1:", *pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("Value 1:", *pointer1)
	fmt.Println("Value 1:", value1)
}
