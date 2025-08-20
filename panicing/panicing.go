package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func getMessage() string {
	fmt.Println("getting message!")
	return "This is a message"
}

func mightFail() {
	fmt.Println("in mightFail")
	if rand.Intn(3) > 1 {
		panic(errors.New("this broke"))
	}
}

func operation() {
	defer func(msg string) {
		fmt.Println("executing deferred function", msg)
		p := recover()
		if p != nil {
			fmt.Println("Recovering from", p)
		} else {
			fmt.Println("it all worked properly")
		}
	}(getMessage())

	fmt.Println("executing body of operation function")
	mightFail()
	if rand.Intn(3) > 1 {
		fmt.Println("leaving early")
		return
	}
	fmt.Println("still going")
}

func main() {
	operation()
	fmt.Println("Carrying on, chaps!")
}
