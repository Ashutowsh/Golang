package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang.")
	greeter()

	ans := adder(4, 6)
	fmt.Println(ans)

	proAns := proAdder(10, 20, 100, 11)
	fmt.Println(proAns)

	result, str := solve(10, 20)

	fmt.Println("Result : ", result, " \nString : ", str)
}

func greeter() {
	fmt.Println("Namastey from golang.")
}

func adder(value1 int, value2 int) int {
	return value1 + value2
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}

	return total
}

func solve(val1 int, val2 int) (int, string) {
	return adder(val1, val2), "Your Answer"
}
