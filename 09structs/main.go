package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs")

	//no inheritance in golang , no super or parent

	tj := User{"TJ", 29}

	fmt.Println(tj)
	fmt.Println(tj.Age)
}

type User struct {
	Name string
	Age  int
}
