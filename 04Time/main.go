package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time package")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// this 01-02-2006 format is constant, this is the reference go follows
	// changing this will not work
	// follow the documentation for more
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	// year, time.month, date, hour, miute, second, nano second, location
	createDate := time.Date(2025, time.August, 9, 5, 5, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))

}
