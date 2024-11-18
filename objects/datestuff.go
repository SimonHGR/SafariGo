package main

import (
	"fmt"
	"objects/date"
)

func main() {
	// var today Date
	// today := Date{}
	// today := Date{Month: 11, Day: 18, Year: 2024}
	// today := date.Date{Month: 11, Day: 18} // no "constructors"
	today := date.MakeDate(18, 11, 2024) // but can encapsulate and make "factories"
	fmt.Printf("today is %v, of type %T\n", today, today)
	// fmt.Printf("day is %d\n", date.GetDay(today)) // uses the "wrong way" function (and works..)
	fmt.Printf("day is %d\n", today.GetDay())
	// first form is "logical" and legal
	// (&today).SetDay(30)
	// second form is "convenient" -- this form is potentially confusing
	today.SetDay(30)
	fmt.Printf("day is %d\n", today.GetDay())

	newYears := date.MakeHoliday(1, 1, 2025, "New Years Day")
	fmt.Println("new Years is", newYears)

	fmt.Printf("day of new years is %d\n", newYears.GetDay())
}
