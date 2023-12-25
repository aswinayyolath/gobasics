package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	message := "Welcome to user input program!"
	fmt.Println(message)
	fmt.Println("==========================")
	reader := bufio.NewReader(os.Stdin)

	// comma ok || err on syntax
	// This is a replacement for try...catch in other languages
	fmt.Print("enter something! : ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered", input)

	// Handling integer input
	fmt.Print("Enter a number: ")
	inputNumber, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	number, _ := strconv.ParseFloat(strings.TrimSpace(inputNumber), 64)

	fmt.Println()
	fmt.Println("Adding one to the input number: ", number+1)

	fmt.Println()

	fmt.Println("Playing with Pointers")
	var ptr *int
	fmt.Println("Value of ptr: ", ptr)

	newNum := 99

	var numPtr = &newNum

	fmt.Println("Memory address value ", numPtr)
	fmt.Println("value at address ", numPtr, "is ", *numPtr)

	*numPtr = *numPtr + 10
	fmt.Println("Now the actual value in the memory is being modified", *numPtr)

	// Array

	var fruitList [5]string
	fruitList[0] = "apple"
	fruitList[3] = "Orange"
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))
	fmt.Printf("Type of fruitList is %T", fruitList)
	fmt.Println()

	// Slices

	var languageList = []string{"Python", "Java", "Ruby"}
	fmt.Printf("Type of languageList %T", languageList)

	fmt.Println()

	// Adding data to slice
	languageList = append(languageList, "Golang", "javaScript")
	fmt.Println(languageList)

	fmt.Println("Slicing language List")
	fmt.Println(languageList[1:]) // Range are always non inclusive
	fmt.Println(languageList[1:3])

	// Creating Slice using make keyword

	highScore := make([]int, 4)
	fmt.Println(highScore) // Output is [0 0 0 0]
	highScore[0] = 100
	highScore[1] = 86
	highScore[2] = 89
	highScore[3] = 91
	// We cannot add highScore[4] = 91 as the default allocation doesn't support this
	// We will get an index out of range error
	// highScore[4] = 12  => not possible

	// But we can use append method to add new element  to the slice

	highScore = append(highScore, 513, 429, 320) // As soon as you use the method append
	// it is going to re allocate the memory and all the values will be accomadated

	fmt.Println(highScore)

	// Sorting an interger slice in ascending order

	sort.Ints(highScore)
	fmt.Println(highScore)

	//Boolean functions in Slice
	//For example if we want to check if the Integers are sorted then run

	fmt.Println(sort.IntsAreSorted(highScore))

	// How to remove a value for slices based on an index

	var courses = []string{"react", "java", "ruby", "python"}
	fmt.Println(courses)

	index := 2

	fmt.Println("===========================")
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

	// Maps in golang

	// Map syntax map[key]value
	subjects := make(map[int]string)
	subjects[0] = "Physics"
	subjects[1] = "Chemistry"
	subjects[2] = "Computer Science"
	subjects[3] = "Biology"
	subjects[4] = "Maths"
	subjects[5] = "English"
	subjects[6] = "Commerce"

	fmt.Println(subjects)

	fmt.Println(subjects[0])

	// Deleting value from a map
	// Syntax delete(map-name, key)

	fmt.Println("Deleting key 1 - chemistry from the map")
	delete(subjects, 1)
	fmt.Println(subjects)

	// Looping through maps

	for key, value := range subjects {
		fmt.Printf("The key is % v and the value is %v\n", key, value)
	}

	// calling Structure
	fmt.Println("Printing User Struct")
	aswin := User{"Aswin", "aswin@gmail.com", true, 26}
	fmt.Println(aswin)

	// Printing Struct details in a more detailed way

	fmt.Println("Printing detailed User details")
	fmt.Printf("More detail of user Aswin %+v", aswin)

	// Printing individual Struct details
	fmt.Println(" Printing individual Struct details")
	fmt.Printf("Name is %v and Email is %v\n\n", aswin.Name, aswin.Email)
	// Golang special case if syntax
	// Where we can initialize a variable
	// and check a condition on that variable

	if num := 10; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is 10 or greater")
	}

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
