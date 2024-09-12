package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to Golang.")
	performGet()
}

func performGet() {
	const myUrl string = "https://ocalhost:8000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Content Length : ", response.ContentLength)

	// content, _ := ioutil.ReadAll(response.Body)

	// fmt.Println(content)
	// fmt.Println(string(content))

	// Another method

	var resString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)

	byteCount, _ := resString.Write(content)
	fmt.Println("ByteCount is : ", byteCount)

	fmt.Println(resString.String())
}
