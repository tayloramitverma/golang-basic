package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang.")

	greetings()

	restul := adder(4, 5)
	fmt.Println("Result is: ", restul)

	proRestul, myMessage := proAdder(4, 5, 1)
	fmt.Println("Result is: ", proRestul)
	fmt.Println("Message is: ", myMessage)

}

func greetings() {
	fmt.Println("Namasty from Golang.")
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(value ...int) (int, string) {
	total := 0

	for _, val := range value {
		total += val
	}

	return total, "Got Sum of enter values"
}
