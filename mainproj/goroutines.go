package main

import (
	"errors"
	"fmt"
)

func produce(out chan int, control chan bool) {
	fmt.Println("starting producer")
	for x := 0; x < 1000; x++ {
		out <- x
	}
	fmt.Println("producer completed")
	control <- true
}

func consume(in chan int, control chan bool) {
	fmt.Println("consumer started")
	for x := 0; ; x++ {
		select {
		case y := <-in:
			if y != x {
				panic(errors.New("uh oh, values received incorrectly!"))
			}
		case <-control:
			fmt.Println("control requests shutdown")
			break
		}
	}
	fmt.Println("consumer finished")
}

func main() {
	comms := make(chan int)
	control := make(chan bool)
	shutdown := make(chan bool)
	go produce(comms, control)
	go consume(comms, shutdown)

	<-control
	shutdown <- true
	for x := 0; x < 1_000_000_000; x++ {
	}
	fmt.Println("shutting down")
}
