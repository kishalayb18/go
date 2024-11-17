package main

import "fmt"

func main() {
	fmt.Println("struct")

	// go does not have classes
	// go has struct similar to classes

	// there is no concept of inheritance in go
	// there is no super class or parent class

	kishalay := User{"Kishalay", "kishalay@google.com", true, 98756}
	fmt.Println("user data ", kishalay)
	// detailed structure data
	fmt.Printf("user data %+v\n", kishalay)
	//print particular fileds
	fmt.Printf("user name here %v \nmail %v\n", kishalay.Name, kishalay.Mail)
}

// name of the structure here User
// name of the structure should be uppercase
type User struct {

	// here the variable needs to PUBLIC
	// to make the variable public variable should start with the uppercase
	Name   string
	Mail   string
	Status bool
	Phone  int
}
