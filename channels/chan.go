package main

import "fmt"

func prod(data chan int, control chan bool) {
	fmt.Println("starting producer")
	for i := 0; i < 1000; i++ {
		if i == 500 {
			data <- -100
		} else {
			data <- i
		}
	}
	fmt.Println("completed producer")
	control <- true
}

func cons(data chan int, control chan bool) {
	fmt.Println("consumer starting")
	// SEE ALSO SELECT!!!!
	for i := 0; i < 1000; i++ {
		v := <-data
		if v != i {
			fmt.Println("Error!", i)
		}
	}
	fmt.Println("consumer completed")
	control <- true
}

func main() {
	data := make(chan int)
	control := make(chan bool)
	go prod(data, control)
	go cons(data, control)
	<-control
	<-control
	// x := <-control
	// y := <-control
	// fmt.Println("x", x, "y", y)
}
