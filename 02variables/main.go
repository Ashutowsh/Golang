package main

import "fmt"

const LoginToken string = "sfvkafbkabfkafbakf"

// capital accessible anywhere

func main() {
	// var username string = "Ashutosh"
	// fmt.Println(username)
	// fmt.Printf("Variable is of type : %T \n", username)

	// var isWorking bool = false
	// fmt.Println(isWorking)
	// fmt.Printf("Variable is of type : %T \n", isWorking)

	// var smallValue uint8 = 15
	// fmt.Println(smallValue)
	// fmt.Printf("Variable is of type : %T \n", smallValue)

	// var a int = 15
	// fmt.Println(a)
	// fmt.Printf("Variable is of type : %T \n", a)

	// var b float64 = 15222.1234567654334565436543345
	// fmt.Println(b)
	// fmt.Printf("Variable is of type : %T \n", b)

	// var str bool
	// fmt.Println("String : ", str)

	// implicit

	var hello = "hello World"
	fmt.Println(hello)

	// no var style  ---> Do not used outside the method

	nUser := 300000
	fmt.Println(nUser)

	fmt.Println("Login Token : ", LoginToken)
}
