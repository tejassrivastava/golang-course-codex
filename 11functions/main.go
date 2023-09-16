package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions")
	user := User{"a", 1}
	fmt.Println(user)
	user.GetName()
	user.SetName()
	user.GetName()
	fmt.Println(user.Name)
}

type User struct {
	Name string
	Age  int
}

func (u User) GetName() {
	fmt.Println("Name :: ", u.Name)
}

func (u User) SetName() {
	fmt.Println("In SetName")
	u.Name = "B"
	fmt.Println("Set Name , Name", u.Name)
}
