package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slices")

	// size is dynamic here

	// how to initialize
	var sl = []string{"kishalay", "gcp"}
	// fmt.Println("type of sl %T ", sl)
	fmt.Println("print sl ", sl)

	// add new values with append method
	sl = append(sl, "cg", "hy", "frd")
	fmt.Println("print new sl ", sl)

	//have the [start:range] here
	//sl[1:4] will print second, third, fourth // fifth sl[4] will not be there
	//sl[:4] will print till the forth value from start
	// sl[3:] will print from sl[3] to the rest of the values
	fmt.Println("[start: range] start=1, range=4 will print 1,2,3 \n", sl[1:4])

	// make keyword
	highScores := make([]int, 5)
	fmt.Println("print default lenght ", len(highScores))

	highScores[0] = 95
	highScores[1] = 70
	highScores[2] = 85
	highScores[3] = 97
	highScores[4] = 72

	//this will throw outofbound as the default memory is size is 5
	// highScores[5] = 75
	fmt.Println("print higgscores ", highScores)

	// but with append method the memory allocation will become dynamic here
	// this will be the benefit over array
	highScores = append(highScores, 75, 97, 83)

	fmt.Println("print new higgscores ", highScores)
	fmt.Println("print new lenght ", len(highScores))

	// sort
	fmt.Println("if sorted ", sort.IntsAreSorted(highScores)) // this returns boolean value
	sort.Ints(highScores)
	fmt.Println("sorted ", highScores) // this will sort the list
	fmt.Println("if sorted ", sort.IntsAreSorted(highScores))

	// remove a value from a particular position
	var nameList = []string{"john", "steve", "jane", "potter", "swap", "sp"}
	fmt.Println("before removing ", nameList)
	// want to remove nameList[2]
	var i int = 2
	// the three dots is required
	nameList = append(nameList[:i], nameList[i+1:]...)
	fmt.Println("removed from position 2 \nnew list ", nameList)

}
