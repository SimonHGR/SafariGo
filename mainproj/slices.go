package main

import "fmt"

func main() {
	n1 := [...]int{1, 2, 3}
	var n2 [3]int
	n2 = n1
	fmt.Println(n1, n2)
	// can compare arrays (of comparable elements) provided
	// element count is the same
	fmt.Println(n1 == n2)

	nums := []int{0, 1, 2, 3, 4, 5, 6}
	//nums2 := []int{0, 1, 2, 3, 4, 5, 6}
	//fmt.Println(nums == nums2) // cannot compare Slices
	fmt.Printf("nums is %v, type is %T\n", nums, nums)

	numsPart := nums[2:5]
	fmt.Println(numsPart)
	numsPart[1] = 99
	fmt.Println(numsPart)
	fmt.Println(nums)

	numsPart = append(numsPart, 15, 16, 17, 18)
	fmt.Println(numsPart)
	fmt.Println(nums)

	for i, v := range numsPart {
		fmt.Printf("index %d, value %d address %p\n", i, v, &v)
		v += 1000
		fmt.Println("index", i, "value", v)
		numsPart[i] -= 1000
		fmt.Println("index", i, "value", v)
	}

	fmt.Println(numsPart)
}
