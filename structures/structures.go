package main

import "fmt"

type Person struct {
	Name, Address string
	Credit        int
}

type Employee struct {
	Person
	Job string
}

func (p Person) getLable() string {
	return fmt.Sprintf("To %s at %s\n", p.Name, p.Address)
}

func increaseCredit(p *Person) {
	fmt.Println("p is ", p)
	p.Credit += 10000
	fmt.Println("p is ", p)
}

func main() {
	fred := Person{Address: "Over here", Name: "Fred", Credit: 1000}
	fmt.Println(fred)

	fmt.Println(fred.getLable())

	barney := Employee{Person{"Barney", "over there", 2000}, "Welder"}
	fmt.Println(barney)
	fmt.Println("barney's name is ", barney.Name)
	fmt.Println(barney.getLable())

	fmt.Println("fred is ", fred)
	increaseCredit(&fred)
	fmt.Println("fred is ", fred)

}
