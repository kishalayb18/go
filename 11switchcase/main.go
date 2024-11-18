package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch case")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("open")
	case 2:
		fmt.Println("move two")
	case 3:
		fmt.Println("move three")
		// fallthrough meanss both case 3, case 4 will work here
		fallthrough
	case 4:
		fmt.Println("move four")
	case 5:
		fmt.Println("move five")
	case 6:
		fmt.Println("move 6")
	default:
		fmt.Println("wrong dice value")
	}
}
