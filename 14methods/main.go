package main

import "fmt"

func main() {
	fmt.Println("Structs ")

	// no inheritence in golang; No super or parent

	ashutosh := User{"Ashutosh", "ashu@go.dev", true, 20}

	fmt.Println(ashutosh)

	fmt.Printf("Ashutosh Details are : %+v\n", ashutosh)

	fmt.Printf("Name is %v and email is %v", ashutosh.Name, ashutosh.Email)
	ashutosh.GetStatus()

	ashutosh.NewEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (user User) GetStatus() {
	fmt.Println("\n\nIs user active : ", user.Status)
}

func (user User) NewEmail() {
	user.Email = "ashu1@go.dev"
	fmt.Println("Eail of this user is : ", user.Email)
}
