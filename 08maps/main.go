package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Map")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["JV"] = "Java"

	fmt.Println(languages)
	fmt.Println(languages["JS"])
	fmt.Println(len(languages))

	delete(languages, "JS")
	fmt.Println(languages)

	word := map[string]int{"a": 1, "b": 2}
	fmt.Println(word)

	for k, v := range word {
		fmt.Printf("%v,%v \n", k, v)
	}

}
