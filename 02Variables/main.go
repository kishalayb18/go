package main

import "fmt"

// this is a global variable
var globalJwt = 30153689

// constant variable
// this is a Public variable as the variable starts with Uppercase
const HyLog = "fjkshutyfsf"

func main() {
	var username string = "kishalay"
	fmt.Println(username)
	fmt.Println("variable type : %T \n", username)

	var isLog bool = true
	fmt.Println(isLog)
	fmt.Println("variable type: %T \n", isLog)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Println("variable type: %T \n", smallValue)

	var smallFloat float32 = 255.370155729
	fmt.Println(smallFloat)
	fmt.Println("variable type: %T \n", smallFloat)

	//implicit type
	var website = "google"
	fmt.Println(website)

	// no var style
	// this is not allowed globally
	noLog := 300
	fmt.Println(noLog)

	//print global variable
	fmt.Println(globalJwt)

	//print constant variable
	fmt.Println(HyLog)

}
