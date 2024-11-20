package main

import "fmt"

func main() {
	fmt.Println("methods")

	kishalay := User{"Kishalay", "kishalay@google.com", true, 98756}
	fmt.Println("user data ", kishalay)
	fmt.Printf("user data %+v\n", kishalay)
	fmt.Printf("user name here %v \nmail %v\n", kishalay.Name, kishalay.Mail)

	// method call
	kishalay.GetStatus()
	fmt.Println("previous mail = ", kishalay.Mail)
	kishalay.newMail()
	// the newMail() method does not change the original value of the mail
	// it just creates a copy of the structure and make changes on that
	// to resolve this we have pointers
	fmt.Println("modified mail = ", kishalay.Mail)
}

type User struct {
	// here the variable needs to PUBLIC
	// to make the variable public variable should start with the uppercase
	Name   string
	Mail   string
	Status bool
	Phone  int
}

// method definition
// method name should start with uppercase if we want to make that public
// we need to pass the structure object to build the method
func (u User) GetStatus() {
	fmt.Println("Status = ", u.Status)
}

// modify the mail
func (u User) newMail() {
	u.Mail = "newmail.com"
	fmt.Println("new mail = ", u.Mail)
}
