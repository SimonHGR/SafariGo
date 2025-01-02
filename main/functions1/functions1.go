package main

import "fmt"

// func add(a int, b int) int {
// func add(a, b int) int {
// 		return a + b
// }

func add(a, b int) (sum int) {
	sum = a + b
	// return sum
	return
}

// report error if a or b are not positive
func addPositives(a, b int) (sum int, ok bool) {
	if a < 0 || b < 0 {
		return 0, false
	} else {
		return a + b, true
	}
}

func main() {
	fmt.Println("xxx")
	fmt.Println("2 + 3 is", add(2, 3))

	// var sum int
	// var ok bool
	// sum, ok = 99, false

	// sum, ok := addPositives(1, -3)
	// fmt.Println("sum is", sum, "ok is", ok)

	if sum, ok := addPositives(1, -3); ok {
		fmt.Println("sum is", sum, "ok is", ok)
	} else {
		fmt.Println("bad input values")
	}
	// fmt.Println(sum) // out of scope!!!

	// sum1, ok1 := addPositives(1, 2)
	// fmt.Println(sum1, ok1)

	// underscore allows me to ignore a assignable value.
	sum1, _ := addPositives(1, 2)
	fmt.Println(sum1)

	names := make(map[string]string)
	names["Ayo"] = "On the hill"
	names["Inaya"] = "next door"
	fmt.Println(names["Inaya"])
	huaAddress, ok := names["Hua"]
	fmt.Println(huaAddress, ok)

	var addition func(int, int) int
	addition = add

	fmt.Println("sum of 1, 7 is", addition(1, 7))

	// anonymous function
	// var add2 = func(a, b int) int {
	// 	return a + b
	// }

	var maker = func(f func(int, int) int, v int) func(int) int {
		return func(a int) int {
			return f(a, v)
		}
	}

	addOne := maker(add, 1)
	fmt.Println("addOne to 99 is", addOne(99))
}
