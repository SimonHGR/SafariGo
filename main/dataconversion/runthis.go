package main

import (
	"fmt"
	"math/rand"
)

func getNum() int {
	fmt.Println("getting number")
	return rand.Intn(3)
}

func main() {
	fmt.Println("stanalone")

	var num int8 = 99
	var num2 int32 = 10
	// num2 = num // not allowed!
	num2 = int32(num)

	fmt.Printf("num2 is %d, of type %T, and has form: %v\n", num2, num2, num2)

	// switch, simple enough

	switch x := getNum(); x {
	case 0, 1: // commas for alternation.
		fmt.Println("less than 2!")
		// break exists but is NOT required
		// if you WANT to "fall through" you must say so:
		// fallthrough
	default:
		fmt.Println("two or more")
	}

	fmt.Println("-------------------")
	var value any
	value = "Hello"
	// value = 99

	v, ok := value.(int)
	fmt.Printf("v is %v, of type %T, ok is %t\n", v, v, ok)

	switch v2 := value.(type) {
	case int:
		fmt.Println("It's an int, value is ", v2)
	case string:
		fmt.Println("It's a string:", value, "length is", len(v2))
	}
}
