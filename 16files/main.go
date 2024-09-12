package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files.")

	content := "Hello World. I am learning golang."

	file, err := os.Create("./learnGolang.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length : ", length)
	defer file.Close()

	readFile("./learnGolang.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)

	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	fmt.Println("dataByte : ", string(dataByte))
	fmt.Println("dataByte : ", dataByte)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
