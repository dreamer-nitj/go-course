package main

import "fmt"

func main() {
	sum := addAllValues(12, 3, 45, 56)
	fmt.Println(sum)
}

func addAllValues(values ...int) int {
	sum := 0

	for i := range values {
		sum += values[i]
	}
	fmt.Printf("%T\n", values)
	return sum
}
