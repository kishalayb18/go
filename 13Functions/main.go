package main

import (
	"fmt"
)

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
	fmt.Printf("result = %v, \nmessaage = %v\n", newRes, newMsg)
	fmt.Println("\n")

	// variadic function with two returns
	fmt.Println("variadic function with two returns")
	name, result := proResult("kishalay", 75, 72, 95, 91, 97)
	fmt.Printf("Result\n%v total marks %v\n", name, result)
	fmt.Println("\n")

	// if we do not need both the returns, we might skip with _
	// here the msg string will be skipped
	// this is comma, ok
	fmt.Println("variadic funtion")
	res, _ := newSum(10, 58)
	fmt.Println("res = ", res)
	fmt.Println("\n")

	// resurcive funtion factorial
	fmt.Println("recursive funtion")
	f := factorial(5)
	fmt.Printf("result of factorial 5 = %v\n", f)
	fmt.Println("\n")

	// annonymous function
	// function that is declared without any name identifier to refer
	// these functions take inputs and provide outputs like normal functions
	// these functions leveraged for containing functionality that need not be named and for short term usage

	// here the function is geting stored inside a variable
	fmt.Println("annonymous function stored in a variable")
	// there is no name of the function in the function signature
	myFunc := func(a int, b int) int {
		res := a + b
		return res
	}
	fmt.Printf("type of the variable containing the function %T\n", myFunc)
	// caling the variable is similar to calling the function here
	fmt.Println("result = ", myFunc(10, 20))

	// another way to declare this
	var (
		my_func = func(s string) {
			fmt.Println("Hey there,", s)
		}
	)
	my_func("joe")

	fmt.Println("\n")

	// in this annonymous function we do not need to store the function inside a variable
	// we will directly call the function with the parameters
	fmt.Println("annonymous function direct call")
	myFuncResult := func(a int, b int) int {
		res := a + b + 5
		return res
	}(10, 60)

	// here the type of myFuncResult will be int as this will store the return of the function
	fmt.Printf("type of myFuncResult %T\n", myFuncResult)
	fmt.Printf("value of myFuncResult %v\n", myFuncResult)
	fmt.Println("\n")

	// high order function
	fmt.Println("High order functions")
	var choice int
	var radius float64
	fmt.Printf("enter the radius ")
	fmt.Scanf("%f", &radius)
	fmt.Println("Choice\n1. calculate area 2. calculate perimeter")
	fmt.Printf("choice = ")
	fmt.Scanf("%d", &choice)

	if choice == 1 {
		area := calculateArea(radius)
		fmt.Printf("area = %v\n", area)
	} else if choice == 2 {
		perimeter := calculatePerimeter(radius)
		fmt.Printf("perimeter = %v\n", perimeter)
	} else {
		fmt.Println("wrong choice")
	}

}

// custom function definition
// function has three things
// function signature return type, function name, function parameters
func hi() {
	fmt.Println("hi from go")
}

// function with return, parametrs
func adder(a int, b int) int {
	var res = a + b
	return res
}

// Variadic Function : Function that accepts variable number of arguments
// to declare a variadic function the type of the final parameter should be preceded by an ellipsis (three dots)
// the variadic parameter should always be placed towards the end
// the parameters will be of a same data type
// name of the parameter, three dots, type of the parameter
// the parametes will store in the slice here
// we are looping the slice and fetching the values
// comma, ok format is getting used for the for loop
func proSum(values ...int) int {
	sum := 0
	// _ is the index of the slice
	for _, val := range values {
		sum += val
	}
	return sum
}

// we are returning two values here
// we are specifying the variable type as well as variable names that we will return here
// the variables are getting declared in the function signature
// in the return statement in the function body we just need to type return
// it will automatically return the two variables values mention in the signature
func proResult(name string, marks ...int) (studentName string, result int) {
	studentName = name
	fmt.Println("Name of the student : ", studentName)
	result = 0
	for _, number := range marks {
		result += result + number
	}
	return
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

// recursive function
func factorial(n int) int {
	// every factorial function should have a base case
	if n == 1 {
		return 1
	}
	factorialResult := n * factorial(n-1)
	return factorialResult
}

// high order functions
// the functions that receive a function as a parameter or return a function as op
// we leverage high order function for composition, modularity and better code readability

// by this we can create smaller function for a certain logic
// composing a complicated function by passing the smaller functions as parameters

// smaller function blocks
func calculateArea(r float64) float64 {
	area := 3.14 * r * r
	return area
}

func calculatePerimeter(r float64) float64 {
	perimeter := 2 * 3.14 * r
	return perimeter
}
