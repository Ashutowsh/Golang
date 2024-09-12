package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter your name : ")

	// comma    ok syntax || Error syntax

	input, _ := reader.ReadString('\n')
	// 1. _, _ := reader.Read..
	// 2. _, err := ...
	fmt.Println("Hello, ", input)
}
