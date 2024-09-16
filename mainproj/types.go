package main

import (
	"dancingcloudservices.com/firstexperiment/util"
	"fmt"
)

const PI = 3.1415926

/*const (
	//ONE = 1
	//ONEAGAIN
	ZERO = iota
	ONE
	TWO
	THREE
)*/

const (
	ONE = 1 << iota
	TWO
	FOUR
	EIGHT
)

func main() {
	fmt.Println("Hello from types")
	fmt.Println(util.Message)

	var count int64 = 0
	fmt.Println(count)
	var c1 int32
	fmt.Println(c1)
	//c1 = count // Simon's platform defaults to int64
	//count = c1 // still not permitted!??
	count = int64(c1) // cast-like

	// constants have "kind", variables have "type"
	// type is rigid, kind is not
	// constant expressions are calculated by the compiler to very high precision
	//var c2 float64 = 99
	var c2 int64 = 123456789012345678901234 / 12345678901234567890
	fmt.Println(c2)

	// do not have to use var for local variables
	// colon-equals declares and initializes a new local variable
	name := "Ayo"
	fmt.Println(name)
	{
		name := "Ishan"
		fmt.Println(name)
	}
	fmt.Println(name)

	// backticks for raw strings
	greeting := "Hello World"
	fmt.Printf("Part 2 is %s\n", greeting[6:]) // not really "substring", but slice
	fmt.Printf("Part 2 is %s\n", greeting[:5])
	fmt.Println(len(greeting))

	greet2 := greeting // assignment is a COPY of value, not of "reference"
	fmt.Println(greet2)
	// & extracts the storage location/ address
	fmt.Printf("greeting is %s and is at %p\n", greeting, &greeting)
	fmt.Printf("greet2 is %s and is at %p\n", greet2, &greet2)
}
