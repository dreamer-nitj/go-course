package main

import "fmt"

func main() {
	n1, l1 := fullName("zaphod", "beeblebrox")
	fmt.Printf("FullName: %v, Number of chars: %v\n", n1, l1)

	n2, l2 := fullNameNakedReturn("zaphod", "beeblebrox")
	fmt.Printf("FullName: %v, Number of chars: %v\n", n2, l2)
}

func fullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

func fullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return
}
