package main

import "fmt"

// main function is the entry of the go program
func main() {
	fmt.Println("functions")
	// function call
	hi()
	fmt.Println("\n")

	fmt.Println("simple function")
	result := adder(3, 7)
	fmt.Printf("result %v\n", result)
	fmt.Println("\n")

	// here the number of parameters wll be dynamic
	fmt.Println("dynamic parameters")
	fmt.Println(proSum(1, 2, 3, 4, 5))
	fmt.Println(proSum(10, 20, 30))
	fmt.Println("\n")

	// two returns
	// values should be stored in comma, ok method
	fmt.Println("two returns")
	newRes, newMsg := newSum(57, 91)
	fmt.Printf("result = %v, \nmessaage = %v", newRes, newMsg)

	// if we do not need both the returns, we might skip with _
	// here the msg string will be skipped
	// this is comma, ok
	res, _ := newSum(10, 58)
	fmt.Println("res = ", res)

}

// custom function definition
// function has three things
// return type, function name, function parameters
func hi() {
	fmt.Println("hi from go")
}

// annonymous function

// function with return, parametrs
func adder(a int, b int) int {
	var res = a + b
	return res
}

// dynamic parameters
// the parameters will be of a specific type
// name of the parameter, three dots, type of the parameter
// the parametes will store in the slice here
// we are looping the slice and fetching the values
// comma, ok format is getting used for the for loop
func proSum(values ...int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum
}

// function with two return
// here the function will return two values
// this might be of same typ of different type
// as there are multiple return, has to put them in parenthisis > (int, string)
// during the return will be returned as comma seperated variables
// during the function call vaules should be stored in comma, ok method
func newSum(a int, b int) (int, string) {
	fmt.Println("multiple return")
	var res = a + b
	msg := "this is a message from newsum"
	return res, msg
}
