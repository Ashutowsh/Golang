package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch")

	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(6) + 1
	fmt.Println("Value a : ", a)
	switch a {
	case 1:
		fmt.Println("Hello 1")
	case 2:
		fmt.Println("Hello 2")
	case 3:
		fmt.Println("Hello 3")
		fallthrough
	case 4:
		fmt.Println("Hello 4")
	case 5:
		fmt.Println("Hello 5")
	case 6:
		fmt.Println("Hello 6")
	}
}
