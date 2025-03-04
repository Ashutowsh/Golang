package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://wynk.in/music/song/har-har-gange/hu_39513341?q=Har+Har+Gange"

func main() {
	fmt.Println("Welcome")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n", response)

	defer response.Body.Close() // caller's responsibility to close the connection.

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println("\n\n\n\nContent1 : ", content)
}
