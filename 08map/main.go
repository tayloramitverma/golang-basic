package main

import "fmt"

func main() {
	fmt.Println("Welcome to map in golang!")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["Ph"] = "PHP"
	languages["GL"] = "GoLang"

	fmt.Println("My Languages are: ", languages)
	fmt.Println("Js value is: ", languages["JS"])

	delete(languages, "Ph")

	fmt.Println("My Languages are: ", languages)

	for key, value := range languages {
		fmt.Printf("For key %v, For Value %v\n", key, value)
	}

}
