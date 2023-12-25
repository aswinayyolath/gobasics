package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File handling in Golang!!!")

	content := "This is the content to write"

	file, err := os.Create("myfile.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("The length is ", length)
	defer file.Close()
	readFile("myfile.txt")
}

func readFile(filename string) {
	dataByte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(dataByte))
}
