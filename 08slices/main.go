package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello World")

	var fruits = []string{"Apple", "Banana"}
	fmt.Printf("Type of fruits : %T\n", fruits)

	fruits = append(fruits, "Mango")
	fmt.Println("Fruits : ", fruits)
	fmt.Println(len(fruits))

	fruits = append(fruits, "Mango", "Watermelon", "Lemon", "Orange")
	fmt.Println("Fruits : ", fruits)
	fmt.Println(len(fruits))

	fmt.Println(fruits[1:5])

	// Another syntax
	numbers := make([]int, 4)

	numbers[0] = 1000
	numbers[1] = 21111
	numbers[2] = 300000
	numbers[3] = 411
	// numbers[4] = 5  ----> Gives Index Out of Bound Error

	numbers = append(numbers, 1, 10, 5) // ---> Reallocates the memory, hence do not give errors
	fmt.Println(numbers)

	sort.Ints(numbers) // Sorting
	fmt.Println(numbers)

	// Removing an element by using indices
	fmt.Println("----------")
	index := 2
	var names = []string{"Ashutosh", "Cristiano", "Messi", "Ronaldinhio", "Neymar", "Haland", "MBappe"}
	fmt.Println(names)

	names = append(names[:index], names[index+1:]...)
	fmt.Println(names)
}
