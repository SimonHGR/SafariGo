package main

import "fmt"

type Textable interface {
	// notice no prefix argument specified
	// BUT THAT IS MANDATORY
	ToText() string
}

type Parent struct {
	s1 string
}

type Child struct {
	Parent
	s2 string
}

func (self Child) ToText() string {
	return "I'm the ToText form of Child, s1=" + self.s1 + "s2=" + self.s2
}

func (self Parent) String() string {
	return "I'm a parent with s1=" + self.s1
}

func main() {
	p := Parent{"parent thing"}
	// c := Child{ "parent part", "child part"}// NOPE, sorry...
	c := Child{Parent{"parent part"}, "child part"} // NOPE, sorry...
	fmt.Println(p)
	fmt.Println(c)

	var t1 Textable = c
	fmt.Println("Text form of t1 is" + t1.ToText())
}
