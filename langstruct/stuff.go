package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num := rand.Intn(4)
	if num := rand.Intn(4); num > 2 { // parens not used, curlies mandatory!
		fmt.Println("big", num)
	} else {
		fmt.Println("small", num)
	}
	fmt.Println("num is", num)
	fmt.Println("-----------------------------")

mySwitch:
	switch n2 := rand.Intn(4); n2 {
	case 0, 1:
		fmt.Println("very small")
		if rand.Intn(2) > 0 {
			break mySwitch
		}
		fmt.Println("still going")
		// fallthrough // do not need break (but it exists)
	case 2, 3:
		fmt.Println("larger")
	}

	fmt.Println("-----------------------------")
	switch { // looks for a true match
	case rand.Intn(4) > 2:
		fmt.Println("A")
	case rand.Intn(4) > 2:
		fmt.Println("B")
	case rand.Intn(4) > 2:
		fmt.Println("C")
	}

	fmt.Println("-----------------------------")
	// var x any = "X"
	var x any = 99
	switch v := x.(type) {
	case string:
		fmt.Printf("v is of type %T, string of length %d\n", v, len(v))
	case int:
		fmt.Printf("v is of type %T, value is %d\n", v, v)
	}
	fmt.Println("-----------------------------")

	for idx := 0; idx < 5; idx++ {
		fmt.Println("idx is", idx)
	}

	idx := 3
	// for ; ; idx-- {
	for {
		fmt.Println(idx)
		// idx--
		if idx == 2 {
			break
		}
	}
	fmt.Println(idx)

}
