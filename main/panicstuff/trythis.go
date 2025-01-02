package main

import (
	"fmt"
	"math/rand"
)

func mightBreak() {
	if rand.Intn(3) == 0 {
		panic("bad random value")
	}
	fmt.Println("mightBreak ending normally")
}

func useMightBreak() {
	defer func(msg string) {
		p := recover()
		if p != nil {
			fmt.Println("panic occurred", p)
		} else {
			fmt.Println("no panic!")
		}
		fmt.Println("in deferred function, message is", msg)
	}("Message for deferred funct")
	mightBreak()
}

func main() {
	fmt.Println("starting")
	useMightBreak()
	fmt.Println("ending")
}
