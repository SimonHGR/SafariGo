package main

import "fmt"

// func add(a, b int) int {
// 	return a + b
// }

// func add(values ...int) (sum int, count int) { // values will be a slice
func add(values ...int) (sum, count int) {
	// var sum int
	// for /*idx,*/ val := range values {
	for _, val := range values {
		// fmt.Println("at index", idx, "value is", val)
		// fmt.Println("value is", val)
		sum += val
	}
	return sum, len(values) // or count = len(values) then return count or just return
	// return
}

func main() {
	s, c := add(1, 2, 9, 8, 7)
	fmt.Println("sum of 1 and 2 is", s, "and count is ", c)
	fmt.Println(add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	var summer func(...int) (int, int)
	summer = add
	fmt.Println(summer(9, 8, 7))

	// local functions must be anonymous
	var addTwo = func(v int) int { 
		s, _ := add(2, v)
		return s
	}
	fmt.Println("10 addTwo is", addTwo(10))
}
