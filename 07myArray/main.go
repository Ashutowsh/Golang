package main

import "fmt"

func main() {
	fmt.Println("Hello, Learn Arrays")

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[3] = "Mango"

	fmt.Println(fruits, " of size : ", +len(fruits))

	var sports = [3]string{"Football", "Cricket", "Basketball"}

	fmt.Println(sports)
}
