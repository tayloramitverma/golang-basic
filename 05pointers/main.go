package main

import (
	"fmt"
)

func main() {

	var welcome string = "Welcome to a class of pointers"
	fmt.Println(welcome)

	myPointer := 54

	var refPointer = &myPointer

	fmt.Println("My pointer value is ", refPointer)
	fmt.Println("My pointer value is ", *refPointer)

}
