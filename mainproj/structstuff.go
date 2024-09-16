package main

import "fmt"

type Month int

type Person struct {
	// caps? accessible outside the package, lower case gives package level encapsulation
	Name, Address string
	Credit        int
}

func showPerson(p Person) {
	fmt.Printf("Person, name = %s, credit= %d\n", p.Name, p.Credit)
}

func (p Person) String() string {
	return "I'm a Person, name is" + p.Name + "and the erst"
}

func (p Person) asText(decorate string) string {
	return decorate +
		fmt.Sprintf("Person, name is %s, address is %s, credit %d", p.Name, p.Address, p.Credit)
}

func (p Person) getAddress() string {
	return p.Name + " lives at " + p.Address
}

type Addressable interface {
	getAddress() string
}

func (p *Person) incCredit(amt int) {
	fmt.Println("p is", p)
	p.Credit += amt
	fmt.Println("p is", p)
}

type SuperBeing struct {
	//p Person // field of type Person refer to "being.p.Name" and similar
	Person
	power string
	//Name  string
}

type Oddball struct {
}

func (odd Oddball) getAddress() string {
	return "No fields, but here I am!"
}

func main() {
	//one := 1
	//var jan Month = one // problem!!!
	var jan Month = 1 // initialize with const, get auto-conversion
	fmt.Printf("jan is %d, of type %T\n", jan, jan)

	ayo := Person{"Ayo", "Down the hill", 10_000}
	var nobody Person
	fmt.Printf("ayo is %v\n", ayo)
	fmt.Println("nobody is", nobody)
	fmt.Println("ayo's name is", ayo.Name)

	showPerson(ayo)

	fmt.Println(ayo.asText("Call me Mr!"))
	fmt.Println("ayo is", ayo)
	//ayo.incCredit(50_000) // this is normal, but it means the following line implictly!!!
	(&ayo).incCredit(50_000)
	fmt.Println("ayo is", ayo)

	// clark IS A SUPERBEING
	//clark := SuperBeing{Person{Address: "Krypton", Name: "Clark Kent", Credit: 1_000_000}, "x-ray vision", "What's this"}
	clark := SuperBeing{Person{Address: "Krypton", Name: "Clark Kent", Credit: 1_000_000}, "x-ray vision"}
	// NOT ALLOWED!!!!
	//var clark Person = SuperBeing{Person{Address: "Krypton", Name: "Clark Kent", Credit: 1_000_000}, "x-ray vision"}

	fmt.Println(clark.asText("Man of Steel"))

	fmt.Println(clark.Name)
	fmt.Println(clark.Person.Name) // valid, sometimes needed for disambiguation

	fmt.Println(ayo)
	fmt.Println(clark)

	var thing interface{}
	fmt.Printf("thing is %v, %T\n", thing, thing)

	var a1 Addressable = ayo

	fmt.Println(a1.getAddress())

	//a1 = Oddball{}
	a1 = ayo
	fmt.Println(a1.getAddress())

	pers, ok := a1.(Person)
	fmt.Println(pers, ok)

	// look at basic switch stuff
	switch v := a1.(type) {
	case Person:
		fmt.Printf("v is %v of type %T\n", v, v)
	default:
		fmt.Println("somethign else")
	}
}
