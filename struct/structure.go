package main

import "fmt"

type Person struct {
	Name, Address string
}

func (p Person) getLabel() string {
	return "Address is " + p.Address
}

type Employee struct {
	Person
	JobRole string
}

type Company struct {
	HQ string
}

func (c Company) getLabel() string {
	return "send mail to " + c.HQ
}

type Addressable interface {
	getLabel() string
}

func main() {
	// ayo := Person{}
	ayo := Person{"Ayo", "Down the lane"}
	// ayo := Person{Address: "Down the lane"}
	// ayo := Person{Address: "Down the lane", Name: "Ayo"}
	// ayo := new(Person) // new makes a pointer!
	// ayo.Name = "Ayo"
	// ayo.Address = "Down the lane"

	fmt.Printf("ayo is %v, Name is %s\n", ayo, ayo.Name)

	// mailPerson := new(Employee)
	mailPerson := Employee{Person{"Ishan", "Next Door"}, "Mail Person"}
	fmt.Printf("mailPerson name %s, address %s, role %s\n",
		mailPerson.Name, mailPerson.Address, mailPerson.JobRole)

	// do not need -> for pointer dereferece
	label := ayo.getLabel()
	fmt.Println("sending mail to", label)
	fmt.Println("mail person, send mail to:" + mailPerson.getLabel())

	megaCorp := Company{"In the Mansion"}
	destinations := []Addressable{mailPerson, megaCorp, ayo}
	for _, a := range destinations {
		fmt.Printf("%v, has address label %s\n", a, a.getLabel())
	}

}
