package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to files in Golang")

	content := "This tes testing content for write new file"

	file, err := os.Create("./myContent.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("Test length is : ", length)

	defer file.Close()

	readFile("./myContent.txt")

}

func readFile(filename string) {
	dataBytes, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("File data is: ", string(dataBytes))
}
