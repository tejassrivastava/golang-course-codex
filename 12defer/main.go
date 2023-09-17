package main

import "fmt"

func main() {
	fmt.Println("Welcome to Defer")
	defer fmt.Println("1")
	fmt.Println("2")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
