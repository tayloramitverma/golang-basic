package main

import (
	"fmt"
	"sort"
)

func main() {

	var welcome string = "Welcome to array in golan"
	fmt.Println(welcome)

	var vegList = []string{"a", "b", "c"}

	vegList = append(vegList, "d", "e")
	fmt.Println("My veg list is: ", vegList)
	vegList = append(vegList[1:3])
	fmt.Println("My veg list is: ", vegList)
	fmt.Println("My veg list is: ", len(vegList))

	highScore := make([]int, 3)

	highScore[0] = 2
	highScore[1] = 5
	highScore[2] = 3
	//highScore[3] = 6

	highScore = append(highScore, 8, 7)

	fmt.Println("My high score is: ", sort.IntsAreSorted(highScore))

	sort.Ints(highScore)

	fmt.Println("My high score is: ", sort.IntsAreSorted(highScore))

	//how to remove values form slices based of index value

	mylangs := []string{"JavaScript", "React", "Next", "Node", "PHP", "Python", "Golan"}

	var index int = 4

	mylangs = append(mylangs[:index], mylangs[index+1:]...)

	fmt.Println("My updated langs are: ", mylangs)

}
