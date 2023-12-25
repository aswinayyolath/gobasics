package main

import "fmt"

func main() {

	// If there is multiple functions which has `defer`ed then the order of execution
	// will be LIFO (Last In First Out), In the below case it will be Three, Two, One
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")

	// This will get executed first as there is no defer keyword
	fmt.Println("Hello")

	deferFunc()
}

func deferFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
