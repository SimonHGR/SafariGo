package main

import (
	"fmt"
	"dancingcloudservices.com/support"
)

// var message = "Hello"
type Month int

func DayOfWeek(d int, m Month, y int) int {
	return 0
}

const (
	ZERO = iota
	ONE
	TWO
	THREE
)

const (
	ONEB = 1 << iota
	TWOB
	FOURB
	EIGHTB
)

func main() {
	// var message string
	// var message string = "Hello"
	// var message = "goodbye"
	message := "hello" // only valid for local variables
	fmt.Println("Hello Go World!")
	fmt.Printf("message is %s, length is %d, equals \"\"? %t\n",
		message, len(message), (message == ""))
	{
		message := "another"
		fmt.Printf("message is %s, length is %d, equals \"\"? %t\n",
			message, len(message), (message == ""))
	}
	var num int
	fmt.Println(num)
	var num2 int64
	num2 = int64(num)
	fmt.Println(num2)

	var m Month = 12
	m = Month(num)
	fmt.Println(m)

	fmt.Printf("ZERO is %d, ONE is %d, %d, %d\n", ZERO, ONE, TWO, THREE)
	fmt.Println(ONEB, TWOB, FOURB, EIGHTB)

	fmt.Println("support message is ", support.SupMsg)
}
