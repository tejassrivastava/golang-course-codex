package main

import "fmt"

func main() {
	fmt.Println("Welcome to If Else")
	var count int = 2
	var user string = "A"

	if count < 3 {
		user = "B"
	}

	fmt.Println(user)

}
