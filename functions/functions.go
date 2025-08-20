package main

import "fmt"

func addUp(v ...int) (sum int) {
	fmt.Printf("v is %T, %v, has %d elements\n", v, v, len(v))
	for _, v1 := range v {
		sum += v1
	}
	return sum
}

func addAndMax(a, b int) (sum, max int) {
	sum = a + b

	if max = a; b > a {
		max = b
	}
	// return
	return sum, max
}

// func add(a, b int) (sum int) {
// 	sum = a + b
// 	// return sum
// 	return
// }

func add(a, b int) int {
	sum := a + b
	return sum
}

// func add(a, b int) int {
// func add(a int, b int) int {
// 	return a + b
// }

func main() {
	fmt.Println("one plus two is", add(1, 2))

	s, m := addAndMax(99, 4)
	fmt.Println("s", s, "m", m)

	fmt.Println("sum of 1..10 is", addUp(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	addTwoNums := func(a, b int) int {
		return a + b
	}

	fmt.Println("adding 1, 2 gives", addTwoNums(1, 2))
	makeAdder := func(a int) func(int) int {
		return func(b int) int {
			return a + b
		}
	}

	addFour := makeAdder(4)
	fmt.Println("addFour(6) is", addFour(6))
}
