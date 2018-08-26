package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "An Implicitly typed string"
	fmt.Println(str1)
	fmt.Printf("str1: %v:%T\n", str1, str1)

	var str2 string = "An explicitly typed string"
	fmt.Printf("str2: %v:%T\n", str2, str2)

	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.Title(str1))

	lValue := "hello"
	uValue := "HELLO"
	fmt.Println("equal?", (lValue == uValue))
	fmt.Println("equal non case sensitive?", strings.EqualFold(lValue, uValue))

	fmt.Println("Contains exp?", strings.Contains(str1, "exp"))
	fmt.Println("Contains exp?", strings.Contains(str2, "exp"))

	// string reader
	reader := strings.NewReader(str1)
	fmt.Println("The number of unread portion of string str1 is:", reader.Len())

	// read one byte at a time
	for reader.Len() != 0 {
		readByte, err := reader.ReadByte()
		if err != nil {
			fmt.Printf("Error reading byte %v\n", err)
		}

		fmt.Printf("readByte:%c and remaining length:%v\n", readByte, reader.Len())
	}

	// lets reset the position
	reader.Reset(str1)
	fmt.Println("the length of unread bytes from str1 after reset:", reader.Len())
}
