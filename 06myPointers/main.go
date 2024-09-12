package main

import "fmt"

func main() {
	fmt.Println("Hello , Let's learn pointers.")

	// var ptr *int
	// fmt.Println("Value : ", ptr)

	num := 23

	var ptr = &num
	fmt.Println("value of actual pointer is ", ptr)
	fmt.Println("Value it is pointing : ", *ptr)

	*ptr = *ptr - 10
	fmt.Println("Value it is pointing : ", *ptr)
}
