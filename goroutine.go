package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex

var signals = []string{}

func main() {

	website := []string{
		"https://google.com",
		"https://fb.com",
		"https://stackoverflow.com",
		"https://github.com",
		"https://medium.com",
		"https://go.dev",
	}

	for _, site := range website {
		go pingSite(site)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

func pingSite(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d is the status code for %s\n", res.StatusCode, endpoint)
}
