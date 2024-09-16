package main

import "fmt"

// func add(x int, y int) int {
//func add(x, y int) int {
//	return x + y
//}

// x is now a slice
//func add(x ...int) int {
//	fmt.Printf("x is %v of type %t\n", x, x)
//	var sum int // init to zero!
//	for _, v := range x {
//		sum += v
//	}
//	return sum
//}

//func add(x ...int) (int, int) {
//	fmt.Printf("x is %v of type %T\n", x, x)
//	var sum int // init to zero!
//	for _, v := range x {
//		sum += v
//	}
//	return sum, len(x)
//}

func add(x ...int) (sum, count int) {
	fmt.Printf("x is %v of type %T\n", x, x)
	for _, v := range x {
		sum += v
	}
	count = len(x)
	return
}

func showSumCount(s, c int) {
	fmt.Println("Sum is", s, "count is ", c)
}

func addTwo(a, b int /*, c float64, d, e int*/) int {
	return a + b
}

func makeIncrementer(amount int) func(int) int {
	return func(y int) int {
		return y + amount
	}
}

func main() {
	//fmt.Println("sum of 1, 2 is", add(1, 2))
	//fmt.Println("sum of 1..10 is", add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	sum, _ := add(1, 2)
	fmt.Println("sum of 1, 2 is", sum)
	var count int
	sum, count = add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("sum of 1..10 is", sum, count)
	fmt.Println(add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	showSumCount(add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	//adder := addTwo // fine, implicit type :)
	var adder func(int, int) int = addTwo // explicit type
	five := adder(2, 3)
	fmt.Println("five is", five)

	// anonymous function
	// functions declared inside other functions MUST be anonymous
	adder = func(a, b int) int {
		fmt.Println("invoking anonymous function")
		return a + b
	}

	seven := adder(4, 3)
	fmt.Println(seven)

	var incBySeven func(int) int = makeIncrementer(7)
	incByNine := makeIncrementer(9)

	fmt.Println("3 in by seven", incBySeven(3))
	fmt.Println("22 in by nine", incByNine(22))
}
