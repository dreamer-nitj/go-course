package main

import (
	"fmt"
	"sort"
)

func main() {
	// keys should be comparable
	states := make(map[string]string)
	fmt.Println(states)

	states["WA"] = "Washington"
	states["OR"] = "Orgaon"
	states["CA"] = "California"
	fmt.Println(states)

	// fetch the value
	california := states["CA"]
	fmt.Println(california)

	// delete a value
	delete(states, "OR")
	fmt.Println(states["OR"])
	fmt.Println(states)

	states["NY"] = "Newyork"

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	fmt.Println("\n Sorted")
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
