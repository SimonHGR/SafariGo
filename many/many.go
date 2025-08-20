package main

import "fmt"

func main() {
	// var names [3]string
	// var names [3]string = [3]string{1: "Fred", 0: "Jim"}
	// var names = [3]string{1: "Fred", 0: "Jim"}
	names := [3]string{1: "Fred", 0: "Jim"}
	fmt.Printf("names is %T, %v, it has %d elements\n", names, names, len(names))
	fmt.Printf("names[0] has length %d\n", len(names[0]))

	// namesSlice := []string{1: "Fred", 0: "Jim"}
	namesSlice := names[0:2]
	fmt.Printf("namesSlice is %T, %v, it has %d elements and cap=%d\n",
		namesSlice, namesSlice, len(namesSlice), cap(namesSlice))

	namesSlice = append(namesSlice, "Inaya", "Ishan")
	fmt.Printf("namesSlice is %T, %v, it has %d elements and cap=%d\n",
		namesSlice, namesSlice, len(namesSlice), cap(namesSlice))
	fmt.Printf("names is %T, %v, it has %d elements\n", names, names, len(names))

}
