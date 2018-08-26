package main

import "fmt"

// Animal interface
type Animal interface {
	Speak() string
}

// Dog struct
type Dog struct {
}

// Speak method
func (d Dog) Speak() string {
	return "Woof!"
}

// Cat struct
type Cat struct {
}

// Speak method
func (c Cat) Speak() string {
	return "Meow!"
}

// Cow struct
type Cow struct {
}

// Speak method
func (c Cow) Speak() string {
	return "Moo!"
}

func main() {
	poodle := Animal(Dog{})
	fmt.Println(poodle)

	animals := []Animal{Dog{}, Cat{}, Cow{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
