package main

import "fmt"

// func add(a int, b int) int {
func add(s string, a, b int) int {
	var c, d int = a, b
	// c = a
	// d = b
	// c, d = a, b
	// c, d := a, b
	fmt.Println("c, d, are", c, d)
	fmt.Println("adding ", a, "and", b)
	return a + b
}

// func addWithExtras(a, b int) (int, bool) {
// 	sum := a + b
// 	isPositive := sum > 0
// 	return sum, isPositive
// }

func addWithExtras(a, b int) (sum int, isPositive bool) {
	sum = a + b
	isPositive = sum > 0
	// return sum, isPositive
	return
}

func showADate(day, month, year int) {
	fmt.Printf("Date is day %d, month %d, year %d", day, month, year)
}

func main() {
	sum := add("", 1, 3)
	fmt.Println("sum is", sum)

	s, p := addWithExtras(9, -14)
	fmt.Println("sum is", s, "positive?", p)

	// NOPE, not "named" argument passing, only positional
	// showADate(month = 11, 18, 2024)
}
