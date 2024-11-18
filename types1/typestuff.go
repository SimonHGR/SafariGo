package main

import "fmt"

type DayNumber int

func numToName(dn DayNumber) string {
	switch dn {
	// switch does not use "break", but break is implicit
	// unless we use "fallthrough"
	case 0:
		return "Saturday"
	case 1:
		return "Sunday"
	default:
		return "Not yet implmented"
	}
}

func main() {
	var dn DayNumber = 1
	fmt.Println(numToName(dn))
	fmt.Println(numToName(0))
	var dn2 int = 1
	// fmt.Println(numToName(dn2)) // NO!!!
	fmt.Println(numToName(DayNumber(dn2))) // OK

}
