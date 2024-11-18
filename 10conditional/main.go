package main

import (
	"fmt"
)

func main() {
	fmt.Println("condition with if else")

	score := 75
	var res string

	if score < 35 {
		res = "not good"
	} else if score < 70 {
		res = "good"
	} else {
		res = "great"
	}
	fmt.Println(res)

	if num := 3; num > 10 {
		fmt.Println("num gt 10")
	} else {
		fmt.Println("num ls 10")
	}

	num1, num2 := 5, 7

	if num1 >= 10 && num2 >= 10 {
		fmt.Println(true)
	} else if num1 < 10 && num2 >= 10 {
		fmt.Println("check num1")
	} else if num1 >= 10 && num2 < 10 {
		fmt.Println("check num2")
	} else {
		fmt.Println(false)
	}

}
