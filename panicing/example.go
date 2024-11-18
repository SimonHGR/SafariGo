package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func getMessage() string {
	fmt.Println("Getting message")
	return "clean up message"
}

func cleanUp(msg string) {
	fmt.Println("Cleaning up, message is", msg)
	if e := recover(); e != nil {
		fmt.Printf("oops, we are recovering from an panic object is %v, of type %T\n", e, e)
	} else {
		fmt.Println("no panic occurred")
	}
}

func mightBreak() string {
	// defer evaluates the argument list immediately
	// but does NOT call the function!!!
	// function is called on the way out of this enclosing function
	defer cleanUp(getMessage())
	if rand.Intn(2) == 1 {
		fmt.Println("mightBreak succeeding")
		return "success"
	}
	fmt.Println("mightBreak failing")
	panic(errors.New("that broke!"))
}

func usesMightBreak() string {
	rv := mightBreak()
	return rv
}

func main() {
	fmt.Println("Starting")
	usesMightBreak()
	fmt.Println("All done")
}
