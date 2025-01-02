package main

import (
	"fmt"
)

type WeddingPhotographer struct {
	Name string
	Fee  int
}

func (p WeddingPhotographer) TakePhoto(subject string) {
	fmt.Println("I'm", p.Name, "taking a picture of", subject, "pay me", p.Fee)
}

type SpySatellite struct {
	Altitude int
	Country  string
}

func (ss SpySatellite) TakePhoto(subject string) {
	fmt.Println("Buzz, click, taking picture of", subject, "from altitude", ss.Altitude)
}

type Photographer interface {
	TakePhoto(s string)
}

// part of the core library -- EVERYTHING satifies this!
// type Any interface {}

func main() {
	var p Photographer
	// p = WeddingPhotographer{"Ayo", 1000}
	// IF SpySatellite TakePhoto required a pointer as the 
	// invocation target (prefix argument), then we would HAVE to say
	// p = &SpySatellite...
	p = SpySatellite{100_000, "NoName"}
	p.TakePhoto("My bridesmaids")
}
