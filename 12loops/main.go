package main

import "fmt"

func main() {
	fmt.Println("Loops break continue goto")

	days := []string{"Sunday", "Monday", "Friday"}
	fmt.Println("\n")

	// for loops
	fmt.Println("simple for loop")
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	fmt.Println("\n")

	// range with for
	fmt.Println("with range")
	for i := range days {
		fmt.Println(days[i])
	}
	fmt.Println("\n")

	// for each loop
	fmt.Println("for each loop")
	for i, day := range days {
		fmt.Printf("for position %v value is %v\n", i, day)
	}
	fmt.Println("\n")

	// similar to while
	fmt.Println("similar to while")
	v1, v2 := 1, 10
	for v1 <= 10 && v2 >= 1 {
		fmt.Printf("values of v1 = %v, v2 = %v\n", v1, v2)
		v1++
		v2--
	}
	fmt.Println("\n")

	// break
	fmt.Println("break")
	v1, v2 = 1, 10
	for v1 <= 10 && v2 >= 1 {
		if v1 > v2 {
			fmt.Printf("\nvalues v1 = %v, v2 =%v, v1>v2 \nbreak loop", v1, v2)
			break
		}
		fmt.Printf("values of v1 = %v, v2 = %v\n", v1, v2)
		v1++
		v2--
	}
	fmt.Println("\n")

	// continue
	fmt.Println("continue")
	v1, v2 = 1, 10

	for v1 <= 10 && v2 >= 1 {
		if v1 == 5 || v2 == 5 {
			fmt.Printf("continue here\n")
			// we need to update the values before we go to continue
			v1++
			v2--
			continue
		}

		fmt.Printf("values of v1 = %v, v2 = %v\n", v1, v2)
		v1++
		v2--
	}
	fmt.Println("\n")

	//goto
	v1 = 0
	for v1 < 10 {
		if v1 == 7 {
			goto newgoto
		}
		fmt.Printf("value now %v\n", v1)
		v1++
	}

	// label for goto
newgoto:
	fmt.Println("this is a label for goto")

}
