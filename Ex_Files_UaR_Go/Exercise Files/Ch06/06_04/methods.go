package main

import "fmt"

// Dog structure
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// Speak the method
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

// SpeakThreeTimes the method, the struct is passed as value so the modifications to
// the Dog variable inside the function won't be visible outside.
func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v! %v! %v!\n", d.Sound, d.Sound, d.Sound)
	fmt.Print(d.Sound)
}

func main() {
	poodle := Dog{"Poodle", 37, "Woof"}
	fmt.Println(poodle)
	poodle.Speak()

	poodle.Sound = "Arf"
	poodle.Speak()

	poodle.SpeakThreeTimes()
	poodle.Speak()
}
