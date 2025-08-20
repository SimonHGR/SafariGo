package main

import (
	"fmt"

	"dancingcloudservices.com/hello/mydata"
	"dancingcloudservices.com/support"
)

func main() {
	fmt.Println("Hello Go World!")
	fmt.Println(mydata.Message)
	// fmt.Println(mydata.whisper)
	fmt.Println(support.SupMessage)
}
