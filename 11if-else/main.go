package main

import "fmt"

func main() {
	if 1 > 2 {
		fmt.Println("Hello")
	} else if 5 < 6 {
		fmt.Println("Else Hello")
	} else {
		fmt.Println("Else Else Hello")
	}
}
