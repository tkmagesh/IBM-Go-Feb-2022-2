package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	resultCh := make(chan int)
	fmt.Println("main started")
	wg.Add(1)
	go add(100, 200, resultCh)
	wg.Done()
	result := <-resultCh
	fmt.Println("result = ", result)
	fmt.Println("main completed")
}

func add(x, y int, ch chan int) {
	fmt.Println("[add] add started")
	time.Sleep(3 * time.Second)
	result := x + y
	fmt.Println("[add] attempting to write result")
	ch <- result
	fmt.Println("[add] completed writing result")
	wg.Done()
}
