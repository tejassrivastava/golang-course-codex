package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files")
	content := "In the file"
	file, err := os.Create("./m.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length is : ", length)
	defer file.Close()
	readFile("./m.txt")
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("File Data : \n", string(data))
}
