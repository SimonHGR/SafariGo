package main

import "fmt"

func changeVals(x, y int) (int, int) {
	return x + 1, y - 1
}

func main() {
	for x := 0; x < 3; x++ {
		fmt.Println("x is", x)
	}

	for x, y := 0, 0; ; {
		fmt.Println("forever", x, y)
		//x++
		//y--
		//x, y = x+1, y-1
		//x, _ = changeVals(x, y) // ignore second value
		x, y = changeVals(x, y)
		if y < -2 {
			break
		}
	}

	//nums := [5]int{1: 1, 0: 2, 2: 3}
	nums := [...]int{1: 1, 0: 2, 2: 3, 7: 99}
	fmt.Println(nums)
	for idx := 0; idx < len(nums); idx++ {
		fmt.Printf("idx = %d, element = %d\n", idx, nums[idx])
	}

	fmt.Println("===============")
	for i, v := range nums {
		fmt.Printf("index %d, value %d\n", i, v)
	}

	word := "英国人"
	for i, v := range word {
		//fmt.Printf("index %d, value %d\n", i, v) // single chars are numbers!
		fmt.Printf("index %d, value %c\n", i, v)
	}
	// as bytes
	for i, v := range []byte(word) {
		fmt.Printf("index %d, value %d\n", i, v) // single chars are numbers!
		//fmt.Printf("index %d, value %c\n", i, v)
	}
}
