package main

import "fmt"

func main() {

	fmt.Println("Welcome to Struct in Golang!")
	// No interitance and No super & parent

	amit := User{"Amit", "av@test.com", true}
	fmt.Printf("Name is %v, Email is %v and Status is %v\n", amit.Name, amit.Email, amit.Status)
	amit.getStatus()
}

// struct

type User struct {
	Name   string
	Email  string
	Status bool
}

// methods

func (u User) getStatus() {
	fmt.Println("User Active status is: ", u.Status)
}
