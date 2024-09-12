package main

import "fmt"

func main() {
	fmt.Println("Hello ")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["Ja"] = "Java"
	languages["Py"] = "Python"
	languages["Go"] = "Golang"

	fmt.Println("List of all languages : ", languages)

	fmt.Println("JS shorts for :", languages["JS"])

	delete(languages, "Py")
	fmt.Println("List of all languages : ", languages)

	for key, value := range languages {
		fmt.Printf("For Key %v value is %v,\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("Values  :  %v,\n", value)
	}
}
