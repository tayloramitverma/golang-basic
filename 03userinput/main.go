package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var welcome string = "Welcome to the user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter rating for my blog:")

	// comma ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you for your rating, ", input)

}
