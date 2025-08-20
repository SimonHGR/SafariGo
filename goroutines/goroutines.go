package main

import "fmt"

func showMessages(control chan bool) {
	for x := 0; x < 10; x++ {
		fmt.Println("Hello...")
	}
	control <- true
}

func main() {
	controlChannel := make(chan bool)
	go showMessages(controlChannel)

	fmt.Println("main still running...")
	// _ = <- controlChannel
	<-controlChannel
	fmt.Println("main shutting down")
}
