package main

import "fmt"

const LoginTokenIs string = "abcdabcdabcd" //public

func main() {
	var username string = "Amit Verma"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isIn bool = true
	fmt.Println(isIn)
	fmt.Printf("Variable is of type: %T \n", isIn)

	var count int = 100
	fmt.Println(count)
	fmt.Printf("Variable is of type: %T \n", count)

	// default values and some aliases

	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n\n", anotherVariable)

	// implicit type
	var website = "beingidea.com"
	fmt.Println(website)

	// no var style
	numberOfUsers := 345000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginTokenIs)

}
