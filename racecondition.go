package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	fmt.Println("Race condition in go!✅")

	var score = []int{0}
	wg.Add(3)
	func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
	}(wg)
	func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
	}(wg)
	func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
	}(wg)

	wg.Wait()
	fmt.Println(score)
}
