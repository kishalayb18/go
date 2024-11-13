package main

import "fmt"

func main() {

	// this is a constant value
	// constant value can not be reassigned further
	const rate = 2

	// for storing values we need not to mention the data type
	var firstValue = 500.1
	var secondValue float64 = 5.7

	// if we just declare a variable we need to mention the variable type
	var thirdValue float64

	// taking the value from the user
	fmt.Print("thirdValue = ")
	fmt.Scan(&thirdValue)

	// another way to declare variable
	// multiple value can be declared in a single line
	// this is the recommended way
	fourthValue, fifthValue := 7, 5

	// for operating two different data types we need to type cast them
	// here we type casted the int to float
	var calculatedValue = float64(firstValue) + secondValue

	// the thirdValue is mentioned as float64 so need not to type cast
	calculatedValue = calculatedValue + thirdValue

	newCalculatedValue := calculatedValue - (float64(fourthValue) + float64(fifthValue))
	newCalculatedValue = newCalculatedValue * rate

	fmt.Println(calculatedValue)
	fmt.Println(newCalculatedValue)

}
