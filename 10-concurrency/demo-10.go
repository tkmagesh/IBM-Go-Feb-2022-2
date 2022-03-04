package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	results := make([]int, 0, 100)
	resultCh := make(chan int)
	fmt.Println("main started")
	for i := 0; i < 100; i++ {
		go add(100, i, resultCh)
	}
	for i := 0; i < 100; i++ {
		result := <-resultCh
		results = append(results, result)
		fmt.Println("result = ", result)
	}
	sort.Ints(results)
	fmt.Println(results)
	fmt.Println("main completed")
}

func add(x, y int, ch chan int) {
	//fmt.Println("[add] add started")
	time.Sleep(3 * time.Second)
	result := x + y
	//fmt.Println("[add] attempting to write result")
	ch <- result
	//fmt.Println("[add] completed writing result")

}
