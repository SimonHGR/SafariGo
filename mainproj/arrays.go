package main

import "fmt"

func addTo(a [3]int) [3]int {
	fmt.Println(a)
	// in Go, assignment, assignment operators, and inc/dec are STATEMENTS
	// not expressions
	a[0] += 1000
	fmt.Println(a)
	return a
}

func addToP(a *[3]int) {
	fmt.Println("in attToP", *a)
	(*a)[0] += 2000
	fmt.Println("in addToP after change", *a)
}

func main() {
	nums := [3]int{
		1,
		2,
		3,
	}
	fmt.Printf("nums is %v at %p\n", nums, &nums)
	nums2 := nums
	fmt.Printf("nums2 is %v at %p\n", nums2, &nums2)
	nums2[0] = 99
	fmt.Printf("nums is %v at %p\n", nums, &nums)
	fmt.Printf("nums2 is %v at %p\n", nums2, &nums2)

	addTo(nums)
	fmt.Println(nums)

	fmt.Println("about to use pointer", nums)
	addToP(&nums)
	fmt.Println("after use of pointer", nums)
}
