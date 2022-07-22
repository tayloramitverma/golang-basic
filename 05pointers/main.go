package main

import (
	"fmt"
)

func main() {
	myPointer := 50
	refPointer := &myPointer

	fmt.Println("My pointer value is: ", myPointer)
	fmt.Println("My pointer memory location is: ", refPointer)
	fmt.Println("My ref pointer value is: ", *refPointer)

	anotherPtr := myPointer
	
	fmt.Println("Another assigned var value is: ", anotherPtr)

	myPointer = 55

	fmt.Println("My pointer value is: ", myPointer)
	fmt.Println("My pointer memory location is: ", refPointer)
	fmt.Println("My ref pointer value is: ", *refPointer)
	fmt.Println("Another assigned var value is: ", anotherPtr)

	*refPointer = *refPointer + 5

	fmt.Println("My pointer value is: ", myPointer)
	fmt.Println("My ref pointer value is: ", *refPointer)
	fmt.Println("Another assigned var value is: ", anotherPtr)
}
