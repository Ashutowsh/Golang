package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://localhost:3000?name=ashutosh&game=FIFA"

func main() {
	fmt.Println("Lean how to handle URL's")
	fmt.Println(myUrl)

	// Parsing the URL

	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("the type of query params are : %T\n", qparams)

	fmt.Println(qparams["name"])

	for _, val := range qparams {
		fmt.Println("Param is : ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "localhost:3000",
		Path:    "/hello",
		RawPath: "user=ashutosh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
