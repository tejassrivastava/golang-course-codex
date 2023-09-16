package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")
	var ptr *int
	fmt.Println(ptr)

	mynumber := 23
	var ptr2 = &mynumber
	fmt.Println(ptr2)
	fmt.Println(*ptr2)
	*ptr2 = *ptr2 + 1

	fmt.Println(*ptr2)
	fmt.Println(mynumber)

}
