package main

import "fmt"

func incrementIt2(vals *[4]int) {
	// vals[0]++
	(*vals)[0]++
	fmt.Printf("%v, at %p\n", *vals, vals)
}

func incrementIt(vals [4]int) {
	vals[0]++
	fmt.Printf("%v, at %p\n", vals, &vals)
}

func main() {
	var nums [4]int
	fmt.Printf("nums is of type %T, value is %v, at %p\n", nums, nums, &nums)
	// var nums2 = [4]int{10, 11, 12, 13}
	var nums2 = [...]int{ 
		10,
		11,
		12,
		13, // trailing either , or }
	}
	// var nums2 = [4]int{0: 10, 3: 13, 1: 11, 2: 12}
	fmt.Printf("nums2 is of type %T, value is %v, at %p\n", nums2, nums2, &nums2)
	nums = nums2 // assignment COPIES and requires same dimension
	fmt.Printf("nums is of type %T, value is %v, at %p\n", nums, nums, &nums)
	nums[0] = 99
	fmt.Printf("nums is of type %T, value is %v, at %p\n", nums, nums, &nums)
	fmt.Printf("nums2 is of type %T, value is %v, at %p\n", nums2, nums2, &nums2)

	incrementIt(nums)
	fmt.Printf("nums is of type %T, value is %v, at %p\n", nums, nums, &nums)

	incrementIt2(&nums)
	fmt.Printf("nums is of type %T, value is %v, at %p\n", nums, nums, &nums)

	// var nums3 [5]int
	// nums3 = nums
	// fmt.Println(nums3)
}
