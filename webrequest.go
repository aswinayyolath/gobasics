package main

import (
	"fmt"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web Request in Golang")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Printf("Response type is %T\n", response)
	fmt.Println(response.Status)
}
