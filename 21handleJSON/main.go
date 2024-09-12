package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name     string
	Username string
	Password string
	Email    string
	Hobbies  []string `json:"Your Hobbies,omitempty"`
}

func main() {
	fmt.Println("Welcome")
	// encodeJSON()
	decodeJSON()
}

func encodeJSON() {

	students := []student{
		{"Ashutosh", "ashu", "12345", "abc@gmail.com", []string{"H1", "H2", "H3"}},
		{"Ashutosh1", "ashu1", "123456", "abc1@gmail.com", []string{"H4", "H5", "H6"}},
		{"Ashutosh2", "ashu2", "1234567", "abc2@gmail.com", []string{"H7", "H8", "H9"}},
		{"Ashutosh3", "ashu3", "12345678", "abc3@gmail.com", nil},
	}

	// package data into JSON

	// finalJSON, err := json.Marshal(students)

	finalJSON, err := json.MarshalIndent(students, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJSON)
}

func decodeJSON() {
	jsonDataFromWeb := []byte(`
	{
        "Name": "Ashutosh",
        "Username": "ashu",
        "Password": "12345",
        "Email": "abc@gmail.com",
        "Your Hobbies": [
                "H1",
                "H2",
                "H3"
        ]
    }
	`)

	var student1 student

	check := json.Valid(jsonDataFromWeb)

	if check {
		fmt.Println("JSON Valid")
		json.Unmarshal(jsonDataFromWeb, &student1)
		fmt.Printf("%#v\n", student1)
	} else {
		fmt.Println("Invalid JSON format")
	}

	// some cases where we want data as key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	fmt.Println("\n")
	for k, v := range myOnlineData {
		fmt.Printf("The key is %v and value us %v and type is %T\n", k, v, v)
	}
}
