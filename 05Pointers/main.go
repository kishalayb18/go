package main

import "fmt"

func main() {
	fmt.Println("pointers")

	// a pointer is the direct reference to the memory address

	// creating a integer type pointer
	// the default value is nil
	//var ptr *int

	myVar := 10
	// referencing the myVar variable and storing the address of that inside newptr
	var newptr = &myVar

	// when we want the value 10, we cal that pointer by reference, *newptr
	var newvar = *newptr

	fmt.Println("value of ptr stores the address of myVar ", newptr)
	fmt.Println("call by reference value of *newptr ", *newptr)
	fmt.Println("value of newvar ", newvar)
}
