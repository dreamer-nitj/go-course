package main

import (
	"fmt"
	"sort"
)

// Person struct
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on the Age field
type ByAge []Person

// Len function
func (a ByAge) Len() int {
	return len(a)
}

// Swap function
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less function
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {
	// without the number between brackets, it becomes a slice
	var colors = []string{"red", "green", "blue"}
	fmt.Println(colors)

	colors = append(colors, "purple")
	fmt.Println(colors)

	// to remove an item from the slice
	colors = append(colors[1:len(colors)]) // this removes the first item from colors
	fmt.Println(colors)

	colors = append(colors[1:])
	fmt.Println(colors)

	// remove the last item
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	// declare the slice a different way
	numbers := make([]int, 5, 5)
	numbers[0] = 1
	numbers[1] = 111
	numbers[2] = 12
	numbers[3] = 11
	numbers[4] = 14
	fmt.Println(numbers)
	fmt.Println("the capacity is:", cap(numbers))

	// now add one more item
	numbers = append(numbers, 98)
	fmt.Println("the capacity is:", cap(numbers))

	// lets sort the slice
	sort.Ints(numbers)
	fmt.Println(numbers)

	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	// There are two ways to sort a slice. First, one can define
	// a set of methods for the slice type, as with ByAge, and
	// call sort.Sort. In this first example we use that technique.
	sort.Sort(ByAge(people))
	fmt.Println(people)

	// The other way is to use sort.Slice with a custom Less
	// function, which can be provided as a closure. In this
	// case no methods are needed. (And if they exist, they
	// are ignored.) Here we re-sort in reverse order: compare
	// the closure with ByAge.Less.
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println(people)
}
