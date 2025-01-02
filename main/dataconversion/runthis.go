package main

import "fmt"

func main() {
	fmt.Println("stanalone")

	var num int8 = 99
	var num2 int32 = 10
	// num2 = num // not allowed!
	num2 = int32(num)

	fmt.Printf("num2 is %d, of type %T, and has form: %v", num2, num2, num2)

}
