package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Welcome to webrequest")
	const urll = "https://lco.dev"
	response, err := http.Get(urll)
	if err != nil {
		panic(err)

	}
	fmt.Println("response is of type :", response.Body)

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("data : ", string(data))

	const myUrl = "https://localhost:3000/api/getUsers?name=t&age=23"

	request, err := response.Request.URL.Parse(myUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println(request.Scheme)
	fmt.Println(request.Host)
	fmt.Println(request.Port())
	fmt.Println(request.Path)
	fmt.Println(request.RawPath)
	fmt.Println(request.RawQuery)
	fmt.Println(request.Query().Get("age"))
	for _, v := range request.Query() {
		fmt.Println(v)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "localhost:3000",
		Path:   "/learn",
	}
	fmt.Println("partsOfUrl : ", partsOfUrl.String())
}
