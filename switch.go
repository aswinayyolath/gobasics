package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("switch and case in golang")
	diceNumber := rand.Intn(6) + 1
	fmt.Println("***", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1, you can Open")
		fallthrough // If this case meets then the next case will also be executed!
	case 2:
		fmt.Println("Dice value is 2, you can move 2 places")
	case 3:
		fmt.Println("Dice value is 3, you can move 3 places")
	case 4:
		fmt.Println("Dice value is 4, you can move 4 places")
	case 5:
		fmt.Println("Dice value is 5, you can move 5 places")
	case 6:
		fmt.Println("Dice value is 6, you can move 6 places")
	}

	days := []string{"Sunday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// Type 1

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// Type 2
	// for day := range days {
	// 	fmt.Println(days[day])
	// }

	// Type 3
	// for i, day := range days {
	// 	fmt.Printf("day on index %v is %v\n", i, day)
	// }

	// if we don't bother about the index then the syntax becomes

	for _, day := range days {
		fmt.Println(day)
	}

	randomvalue := 1

	for randomvalue < 15 {
		if randomvalue == 5 {
			goto someLabel
		}
		fmt.Println(randomvalue)
		randomvalue++
	}

someLabel:
	fmt.Println("Some label")
}
