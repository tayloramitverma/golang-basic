package main

import (
	"fmt"
)

func main() {

	var welcome string = "Welcome to array in golan"
	fmt.Println(welcome)

	var fruitList [5]string

	fruitList[0] = "mango"
	fruitList[1] = "banana"
	fruitList[3] = "apple"

	fmt.Println("My fruit list ", fruitList)
	fmt.Println("My fruit list length", len(fruitList))

	var vegList = [4]string{"a", "b", "c"}
	fmt.Println("My veg list is: ", vegList)
	fmt.Println("My veg list is: ", len(vegList))
}
