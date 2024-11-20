package main

import "fmt"

func main() {
	fmt.Println("defer")

	// defer key word will put the line at the last of the flow
	defer fmt.Println("hello")

	fmt.Println("new hello")

	// defer works with the flow of LIFO
	// so with multiple defere the last input will be the first op
	defer fmt.Println("first")
	defer fmt.Println("second")

	// till here 3 defer statement there

	myDefer()
	// here the defer stack is 0, 1, 2, 3, 4

	// the combined stack
	// hello, first, second, 0, 1, 2, 3, 4

	// the op will be
	// new hello (there is no defer here)
	// 4, 3, 2, 1, 0, second, first, hello (LIFO of the combined stack)

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	// stack will be created
	// 0, 1, 2, 3, 4
}
