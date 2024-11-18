package date

import (
	"errors"
	"fmt"
)

type Date struct {
	// day, month, year int // lower case first letter? only in this package
	day   int
	month int
	year  int
}

type Holiday struct {
	Date
	name string
}

func MakeHoliday(day, month, year int, name string) Holiday {
	return Holiday{Date{day, month, year}, name}
}

func MakeDate(day, month, year int) Date {
	if day < 1 || day > 31 {
		panic(errors.New("Bad date values"))
	}
	return Date{day, month, year}
}

// wrong way!!!
// func GetDay(d Date) int {
// 	return d.day
// }

// right way (well, "object oriented way")
func (d Date) GetDay() int {
	return d.day
}

// ANY function argument (not only a prefix arg!) that wants to permit
// mutation, MUST be a pointer (or a thing that contains a pointer to the
// thing to be mutated)
// func (d Date) SetDay(newDay int) {
// 	fmt.Printf("d at entry to SetDay is %v, newDay is%d\n", d, newDay)
// 	d.day = newDay
// 	fmt.Printf("d at exit from SetDay is %v\n", d)
// }

func (d *Date) SetDay(newDay int) {
	fmt.Printf("d at entry to SetDay is %v, newDay is%d\n", d, newDay)
	d.day = newDay
	fmt.Printf("d at exit from SetDay is %v\n", d)
}
