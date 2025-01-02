package main

import "fmt"

type Date struct {
	Day, Month, Year int
}

type Holiday struct {
	Date // embedded struct
	Name string
}

// USELESS!!! Pass by value (copy!!!)
// func Tomorrow(d Date) {
// 	d.Day++
// }

func Tomorrow(d Date) Date {
	return Date{d.Day + 1, d.Month, d.Year}
}

func TomorrowP(d *Date) {
	d.Day++
}

func (d *Date) TomorrowO() {
	d.Day++
}

func (d Date) String() string {
	return fmt.Sprintf("I'm a date, day is %d, month is %d, year is %d",
		d.Day, d.Month, d.Year)
}

func main() {
	// var today Date
	// today := Date{2, 1, 2025} // day, month, year
	// today := new(Date) // new returns a POINTER to the allocated struct

	today := Date{Month: 1, Day: 2}
	today.Day = 2 // Go uses dot, not arrow for pointer to member access
	fmt.Printf("today is %v, of type %T\n", today, today)
	tomorrow := Tomorrow(today)
	fmt.Printf("today is %v, of type %T\n", today, today)
	fmt.Printf("tomorrow is %v, of type %T\n", tomorrow, tomorrow)
	TomorrowP(&today)
	fmt.Printf("today is %v, of type %T\n", today, today)

	// Go takes the pointer automatically (usually!)
	// -- How do you know what might happen??
	today.TomorrowO()
	fmt.Printf("today is %v, of type %T\n", today, today)

	// newYearsDay := Holiday{}
	// MUST build the Date part "manually" :)
	newYearsDay := Holiday{Date{1, 1, 2025}, "New Year's Day"}
	fmt.Println("newYearsDay is", newYearsDay)
	fmt.Println("newYearsDay.Day is", newYearsDay.Day)
}
