package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Loop stuff")

	// myloop: for {
	for {
		fmt.Println("forever")
		if rand.Intn(8) > 6 {
			// break myloop
			break
		}
	}

	counter := 4
	for counter > 0 {
		fmt.Println("counter", counter)
		counter--
		// increment and decrement are STATEMENTS, not expressions
		// --counter // NOPE
		// nextVal := counter-- // NOPE
		// assignments do NOT have value in Go either...
	}

	for x, y := 0, 10; x < y; x++ {
		fmt.Println(x, y)
		y--
	}

	name := "你好 Alonzo"
	for i, j := range name {
		fmt.Printf("i is %d, j is %c\n", i, j)
	}

	for _, j := range name {
		fmt.Printf("j is %c\n", j)
	}

	var x, y int
	// _, y = 3, 7
	x, y = 3, 7

	fmt.Println(x, y)

}
