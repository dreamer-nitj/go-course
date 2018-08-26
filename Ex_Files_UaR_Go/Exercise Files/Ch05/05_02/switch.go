package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	//dow := rand.Intn(6) + 1
	//fmt.Println("Day", dow)

	result := ""

	switch dow := rand.Intn(6) + 1; dow {
	case 1:
		result = "Its Sunday!"
	case 7:
		result = "Its Saturday!"
	default:
		result = "Its a Weekday!"
	}

	//fmt.Println("Day", dow, ",", result)
	fmt.Println(result)

	x := -42
	switch {
	case x < 0:
		result = "less than zero"
		// 'fallthrough' keyword executes the next case statements as well
	case x == 0:
		result = "equal to zero"
	default:
		result = "greater than zero"
	}

	fmt.Println(result)
}
