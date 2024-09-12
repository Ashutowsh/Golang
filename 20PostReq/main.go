package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Golang.")
	// performGet()
	PerformJSONReq()
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

func PerformJSONReq() {
	const myurl = "https://localhost:8000/post"

	// fake json payload

	requestBody := strings.NewReader(`
	{
	"coursename" : "Let's go with golang.",
	"price" : 0,
	"platform" : "hello"
 	}`)

	res, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := ioutil.ReadAll(requestBody)
	fmt.Println(string(content))
}

func PostFormDataReq() {
	const myurl = "https://localhost:8000/postform"

	// formData

	data := url.Values{}
	data.Add("firstname", "Ashutosh")
	data.Add("lastname", "Phadke")
	data.Add("Roll no", "20")

	res, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}
