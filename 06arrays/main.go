package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays")

	var fruitList [3]string
	fruitList[0] = "a"
	fruitList[1] = "b"
	fruitList[2] = "c"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList = [3]string{"a", "b", "c"}
	fmt.Println(vegList)

}
