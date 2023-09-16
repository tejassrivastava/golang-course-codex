package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("** WELCOME TO OUR PIZZA APP **")
	// fmt.Println("Enter Yout Name:")
	fmt.Println("Enter Pizza Rating 1-5:")

	userInput, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating ", userInput)

	// numRating, err := strconv.ParseFloat(strings.TrimSpace(userInput), 64)
	numRating, err := strconv.ParseInt(strings.TrimSpace(userInput), 0, 0)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(numRating)
	numRating = numRating + 1
	fmt.Println(numRating)
}
