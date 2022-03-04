package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("main started")
	resultCh := add(100, 200)
	result := <-resultCh
	fmt.Println("result = ", result)
	fmt.Println("main completed")
}

func add(x, y int) chan int {
	resultCh := make(chan int)
	go func() {
		fmt.Println("[add] add started")
		time.Sleep(3 * time.Second)
		result := x + y
		fmt.Println("[add] attempting to write result")
		resultCh <- result
		fmt.Println("[add] completed writing result")
	}()
	return resultCh
}
