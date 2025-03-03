package main

import "fmt"

func main() {
	var nums = [10]int{1, 2, 3, 4, 5, 0, 0, 0, 0, 0}
	var three = nums[2:6] // [:x] from start [x:] from x to end, [:] all of it
	fmt.Printf("three is %v, type %T, len %d, cap %d\n",
		three, three, len(three), cap(three))
	three[0] = 100
	fmt.Printf("three is %v, type %T, len %d, cap %d\n",
		three, three, len(three), cap(three))
	fmt.Printf("nums is %v, type %T, len %d, cap %d\n",
		nums, nums, len(nums), cap(nums))

	// uh oh, will this reallocate?
	three = append(three, -1, -2, -3, -4, -5)
	fmt.Printf("three is %v, type %T, len %d, cap %d\n",
		three, three, len(three), cap(three))
	fmt.Printf("nums is %v, type %T, len %d, cap %d\n",
		nums, nums, len(nums), cap(nums))

	// odd := nums[-2:] // NOPE :)

}
