package main

import "fmt"

type Person struct {
	Name, Address string
}

type Employee struct {
	Person
	JobRole string
}

func main() {
	// ayo := Person{}
	// ayo := Person{"Ayo", "Down the lane"}
	// ayo := Person{Address: "Down the lane"}
	// ayo := Person{Address: "Down the lane", Name: "Ayo"}
	ayo := new(Person) // new makes a pointer!
	ayo.Name = "Ayo"
	ayo.Address = "Down the lane"

	fmt.Printf("ayo is %v, Name is %s\n", ayo, ayo.Name)

	// mailPerson := new(Employee)
	mailPerson := Employee{Person{"Ishan", "Next Door"}, "Mail Person"}
	fmt.Printf("mailPerson name %s, address %s, role %s\n",
		mailPerson.Name, mailPerson.Address, mailPerson.JobRole)
}
