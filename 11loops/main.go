package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops in Golang.")

	days := []string{"Mon", "Tue", "Wed", "Thus", "Fir", "Sat", "Sun"}

	fmt.Println("Week days are: ", days)

	for d := 0; d < len(days); d++ {
		fmt.Println("Days is: ", days[d])
	}

	for i := range days {
		fmt.Println("Days is: ", days[i])
	}

	testValue := 1

	for testValue < 10 {

		if testValue == 2 {
			fmt.Println("Test Value: ", testValue)
			testValue++
			continue
		}

		if testValue == 5 {
			goto msg
		}

		fmt.Println("Test Value is: ", testValue)
		testValue++
	}

	//goto

msg:
	fmt.Println("This danger zone alert")

}
