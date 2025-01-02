package main

import (
	"fmt"
)

func producer(data chan int, control chan bool) {
	fmt.Println("Producer starting")
	for count := 0; count < 2000; count++ {
		if count == 500 {
			data <- -1
		} else {
			data <- count
		}
	}
	fmt.Println("Producer ending")
	control <- true
}

func consumer(data chan int, control chan bool) {
	fmt.Println("Consumer starting")

outer:
	for count := 0; ; count++ {
		select {
		case val := <-data:
			if val != count {
				fmt.Println("oops, bogus data received")
			} else {
				fmt.Println(val)
			}
		case <-control:
			break outer
		}
	}
	fmt.Println("Consumer ending")
	control <- true
}

func main() {
	data := make(chan int)
	control := make(chan bool)
	shutdown := make(chan bool)
	fmt.Println("Main starting")
	go producer(data, control)
	go consumer(data, shutdown)
	<-control
	shutdown <- true
	<-shutdown
	fmt.Println("Main ending")

}
