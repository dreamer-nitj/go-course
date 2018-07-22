package main

import "fmt"

// Vertex is the struct
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell labs"] = Vertex{
		40.68, -74.09,
	}

	fmt.Println(m["Bell labs"])
}
