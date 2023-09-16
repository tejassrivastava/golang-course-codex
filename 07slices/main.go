package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var alph = []string{"a", "b"}
	fmt.Println(alph)
	alph = append(alph, "c", "d")
	fmt.Println(alph)
	fmt.Println(append(alph[1:]))
	fmt.Println(append(alph[1:3]))
	fmt.Println(append(alph[:3]))

	// Use make to define a slice

	score := make([]int, 4)

	score[0] = 1
	score[1] = 55
	score[2] = 3
	score[3] = 4

	fmt.Println(score)

	score = append(score, 5, 6)

	fmt.Println(score)
	sort.Ints(score)
	fmt.Println(sort.IntsAreSorted(score))

	//Remove value from slice based on index

	var word = []string{"a", "b", "c", "d", "e"}
	fmt.Println(word)
	var index int = 2
	word = append(word[:index], word[index+1:]...)
	fmt.Println(word)
}
