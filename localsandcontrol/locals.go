package main

import (
	"fmt"
	"math/rand"
)

const (
	NUMBER_TWO = 2
	HUGE       = 1_000_000_000_000_000_000_000_000_000_000
)

const (
	ONE = 1 << iota
	TWO = 1 << iota
	FOUR
	EIGHT
	FOUR_AGAIN = iota
)

func main() {
	// var name string
	// var name = "hello"
	name := "hello"
	{
		// name := "bonjour"
		name = "bonjour"
		fmt.Println("name is ", name, "length is", len(name))
	}
	fmt.Println("name is ", name, "length is", len(name))

	// var sqrtMinusOne complex128
	var count int16 = 32 * HUGE / HUGE
	var c2 int64 = 64
	c2 = int64(count)
	fmt.Println(c2)
	fmt.Println(ONE, TWO, FOUR, EIGHT, FOUR_AGAIN)

	// die := rand.Intn(6) + 1
	// fmt.Println("dice shows", die)
	// if die == 6 {
	// 	fmt.Println("it's a six!")
	// } else {
	// 	fmt.Println("bad luck")
	// }

	die := 99

	// modify outer? (or not...)
	if die = rand.Intn(6) + 1; die == 6 {
		// if die := rand.Intn(6) + 1; die == 6 {
		fmt.Println("it's a six!")
		fmt.Println("dice shows", die)
	} else {
		fmt.Println("bad luck")
		fmt.Println("dice shows", die)
	}
	fmt.Println("dice (outer) shows", die)

	switch die := rand.Intn(6) + 1; die {
	case 1, 2, 3:
		fmt.Println("It's small!", die)
		fallthrough
	case 4, 5:
		fmt.Println("four or five", die)
	default:
		fmt.Println("It's six!!!")
	}

	target := rand.Intn(3)
	// switch true {
	switch {
	case rand.Intn(3) == target:
		fmt.Println("first try")
	case rand.Intn(3) == target:
		fmt.Println("second try")
	case rand.Intn(3) == target:
		fmt.Println("third try")
	case rand.Intn(3) == target:
		fmt.Println("fourth try")
	default:
		fmt.Println("gave up trying")
	}

	var thing any = "Hello"
	s, ok := thing.(string)
	fmt.Println(s, ok)
	f, ok := thing.(float32)
	fmt.Println(f, ok)

	switch v := thing.(type) {
	case int16:
		fmt.Println("v is an int", v)
	case string:
		fmt.Printf("v is a %T of length %d, it's %s\n", v, len(v), v)
	}

	// there is no ?: conditional expression
}
