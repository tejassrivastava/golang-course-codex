package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome To"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Yout Name:")

	userInput, _ := reader.ReadString('\n')

	fmt.Println("Welcome ", userInput)

}
