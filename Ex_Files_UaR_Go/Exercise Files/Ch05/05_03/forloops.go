package main

import "fmt"

func main() {

	sum := 1
	fmt.Println("Sum:", sum)

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	sum = 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("Sum:", sum)

	for index := 0; index < len(colors); index++ {
		fmt.Println(colors[index])
	}

	for index := range colors {
		fmt.Println(colors[index])
	}
}
