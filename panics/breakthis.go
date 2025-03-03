package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func mightBreak() {
	fmt.Println("starting might break")
	if rand.Intn(2) > 0 {
		panic(errors.New("Something broke"))
	}
	fmt.Println("might break finished normally")
}

func getMessage() string {
	fmt.Println("getting message")
	return "deferral message"
}

func deferedFunc(arg string) {
	fmt.Println("in defered func arg is", arg)
	if p := recover(); p != nil {
		fmt.Println("we recovered from", p)
	} else {
		fmt.Println("no panic seen here!")
	}
}

func useMightBreak() {
	defer deferedFunc(getMessage())
	fmt.Println("about to call might break")
	mightBreak()
	fmt.Println("returned from might break")
}

func main() {
	useMightBreak()
	fmt.Println("finishing main normally")
}
