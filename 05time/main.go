package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study in go.")

	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	// "01-02-2006 15:04:05 Monday" - standard string for formatting.

	createdDate := time.Date(2024, time.January, 16, 12, 11, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
