package main

import "fmt"

func main() {
	fmt.Println("get a taste of function!")
	greet()
	result := adder(2, 3)
	fmt.Println("Result = ", result)

	proAddResult := proAdder(2, 3, 4, 5, 6, 7)
	fmt.Println("Pro Addition Result ", proAddResult)

	addPrintRes, _ := addPrint(10, 100)
	fmt.Println(addPrintRes)

	user := User{"Aswin", 1, 27}
	user.getUserData()
}

func greet() {
	fmt.Println("Hello world!")
}

func adder(num1 int, num2 int) int {
	return num1 + num2
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

// Functions with multiple return type

func addPrint(numbers ...int) (int, string) {
	tot := 0
	for _, v := range numbers {
		tot += v
	}
	return tot, "Some message"
}

/// Methods in golang

type User struct {
	Name string
	Id   int
	Age  int
}

func (u User) getUserData() {
	fmt.Println("User Name is ", u.Name)
}
