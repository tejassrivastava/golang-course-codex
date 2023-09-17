package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to API Request")

	getRequest()
	postRequest()
}

func postRequest() {
	//create json payload
	requestBody := strings.NewReader(`
	{"title":"foo",
		"body":"bar",
		"userId": 1 }`)
	fmt.Println("Request Body : ", requestBody)
	response, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json; charset=UTF-8", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Post Response Status :", response.StatusCode)
	fmt.Println("Post Response Body:", response.Body)
	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data : ", string(data))
}

func getRequest() {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code : ", response.StatusCode)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Body :: ", string(body))
	var responseString strings.Builder
	byteCount, err := responseString.Write(body)
	fmt.Println("Byte Count:", byteCount)
	fmt.Println("Body String Len: ", responseString.Len())
	fmt.Println("Body String : ", responseString.String())
}
