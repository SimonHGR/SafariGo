package main

import "fmt"

func addUp(vals ...int) (sum int, count int) {
	fmt.Printf("vals is %v, len is %d\n", vals, len(vals))
	count = len(vals)

	// for idx := 0; idx < len(vals); idx++ {
	// 	sum += vals[idx]
	// }

	// first value is index, second is value from "many"
	for _, v := range vals {
		fmt.Println("adding:", v)
		sum += v
	}

	return sum, count
}

func changeArray(vals ...int) {
	fmt.Printf("vals is %v\n", vals)
	// for _, v := range vals {
	// 	v += 10
	// 	fmt.Printf("value is %d\n", v)
	// }

	for idx, _ := range vals {
		vals[idx] += 10
		fmt.Printf("value is %d\n", vals[idx])
	}

	fmt.Printf("vals is %v\n", vals)
}

func main() {
	fmt.Println("arrays starting")

	// var numbers [4]int = [4]int{1, 3, 5, 7}
	var numbers [4]int = [4]int{
		// 1, 3, 5, 7, // last line ends either , or }
		3: 1, 2: 3, 1: 5, // specifying subscripts (out of order if desired)
	}
	fmt.Printf("numbers is %v, of type %T, length is %d\n",
		numbers, numbers, len(numbers))

	var moreNums [4]int
	moreNums = numbers
	fmt.Printf("moreNums is %v, of type %T, length is %d\n",
		moreNums, moreNums, len(moreNums))

	numbers[0] = 99
	fmt.Printf("numbers is %v, of type %T, length is %d\n",
		numbers, numbers, len(numbers))
	fmt.Printf("moreNums is %v, of type %T, length is %d\n",
		moreNums, moreNums, len(moreNums))

	// var yetMoreNums [5]int
	// yetMoreNums = moreNums // NO! array size is PART OF TYPE
	// need a "slice" -- this is (kinda) a more dynamic array thing...

	var numSlice []int = []int{1, 2, 3, 0, 0, 0}
	numSlice = numSlice[0:3]
	fmt.Printf("numSlice is %v, type %T, len %d, cap %d\n",
		numSlice, numSlice, len(numSlice), cap(numSlice))
	numSlice2 := append(numSlice, 99, 98)
	numSlice[0] = 100
	fmt.Printf("numSlice is %v, type %T, len %d, cap %d\n",
		numSlice, numSlice, len(numSlice), cap(numSlice))
	fmt.Printf("numSlice2 is %v, type %T, len %d, cap %d\n",
		numSlice2, numSlice2, len(numSlice2), cap(numSlice2))

	sum, count := addUp(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("sum of 1 .. 10 is %d, count is %d\n", sum, count)

	changeArray(1, 2, 3, 4)
}
