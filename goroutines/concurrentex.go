package main

import (
	"fmt"
	"time"
)

func counter(control chan<- bool) {
	fmt.Println("Worker starting")
	for idx := 0; idx < 10; idx++ {
		fmt.Println("idx is", idx)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Worker finishing")
	// <- control // would read, but not allowed for write only chan<-
	control <- true
}

func main() {
	control := make(chan bool /*, 10 */)
	fmt.Println("Starting")
	go counter(control)
	fmt.Println("Configured...")
	read := <-control
	fmt.Printf("value read is %v\n", read)
	fmt.Println("Ending")
}
