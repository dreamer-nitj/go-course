package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("filename.ext")

	attendance := map[string]bool{
		"Ann":  true,
		"Mike": true}
	if err != nil {
		fmt.Printf("Error in opening: %v", err)
	}

	fmt.Println(f)

	// create a custom error object
	myError := errors.New("my custome error")
	fmt.Println(myError)

	attended, ok := attendance["Mike"]
	if ok {
		fmt.Println("Mike attended? ", attended)
	} else {
		fmt.Println("No info for mike")
	}
}
