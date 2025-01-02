package main

import "fmt"

func incVal(arr [4]int) {
	// assignment operators, and ++ / -- do NOT have expression value, only side effect
	arr[0] += 10
}

// pointer (perhaps as structure member) is ESSENTIAL if you want to change
// the caller's data!!! (in EVERY situation)
func incValPtr(arr *[4]int) { 
	// assignment operators, and ++ / -- do NOT have expression value, only side effect
	arr[0] += 10
}

func main() {
	// array size is fixed at creation
	// and PART OF THE TYPE
	var nums [4]int
	fmt.Printf("nums is %v, of type %T\n", nums, nums)
	nums2 := [4]int{1, 2}
	fmt.Printf("nums2 is %v, of type %T\n", nums2, nums2)
	nums = nums2 // assignment (including actual parameter passing!!!) is BY VALUE (copy)
	fmt.Printf("nums is %v, of type %T\n", nums, nums)
	nums2[0] = 99
	fmt.Printf("nums is %v, of type %T\n", nums, nums)
	fmt.Printf("nums2 is %v, of type %T\n", nums2, nums2)
	incVal(nums)
	fmt.Printf("nums is %v, of type %T\n", nums, nums)
	incValPtr(&nums) // & gets the address (pointer) to the storage
	fmt.Printf("nums is %v, of type %T\n", nums, nums)

	// moreNums := new([5]int)
	// moreNums = nums // array assignment REQUIRES identical sizes
}
