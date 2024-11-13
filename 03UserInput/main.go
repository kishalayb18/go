package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to Kishalays room"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter number of people for the room ")

	// comma ok
	people, _ := reader.ReadString('\n')
	fmt.Println("number of people stored, ", people)

	//roomSize = people + 1

	// this will not be possible
	// people will be of type string

	numPeople, err := strconv.ParseFloat(strings.TrimSpace(people), 64)
	// if we do not trin the space the input will have a \n character which will not be number

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("new room size ", numPeople+1)
	}
}
