package main

import "fmt"

func main() {
	fmt.Println("arays")

	// here array is homogeneous
	var arr [5]int

	arr[0] = 10
	arr[1] = 20
	arr[3] = 30

	fmt.Println("arr ", arr)

	//lenghth of array
	fmt.Println("lenth of arr ", len(arr))

	var namearr = [5]string{"kishalay", "cg", "hy", "frd", "ggl"}

	fmt.Println("names ", namearr)
}
