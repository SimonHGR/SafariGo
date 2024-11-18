package main

import (
	"fmt" // fmt is standard, so no need to specify a module.

	"dancingcloudservices.com/hello/messages"
)

func main() {
	// println() -- don't use this!
	fmt.Println("Hello Go World!")
	fmt.Println("Other message is", messages.Message)
	// var count int // might want to be more specific, e.g. int64
	// literals and constant expressions are "different"
	// they have "kind" eg, integer, but "type"
	// calculations are done to insane precision
	var count int16 = 10
	var biggerCount int64
	// biggerCount = count // NO, NOT ALLOWED
	biggerCount = int64(count)
	fmt.Printf("biggerCount is %v, type %T\n", biggerCount, biggerCount)

	fmt.Printf("count is %d, %v, %T\n", count, count, count)

	scopeVar := 100
	{
		scopeVar := 10
		fmt.Printf("scopeVar is %d, %T, at %p\n", scopeVar, scopeVar, &scopeVar)
	}
	{
		scopeVar := 10
		fmt.Printf("scopeVar is %d, %T, at %p\n", scopeVar, scopeVar, &scopeVar)
	}
	fmt.Printf("scopeVar is %d, %T, at %p\n", scopeVar, scopeVar, &scopeVar)
	

}
