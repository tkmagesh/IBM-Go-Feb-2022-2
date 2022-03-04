package main

import (
	"fmt"
	"time"
)

func main() {
	count := 20
	ch := genEvenNos(count)
	for i := 0; i < count; i++ {
		evenNo := <-ch
		fmt.Println("even number =", evenNo)
	}
}

func genEvenNos(count int) chan int {
	resultCh := make(chan int)
	go func() {
		for i := 1; i <= count; i++ {
			evenNo := i * 2
			resultCh <- evenNo
			time.Sleep(500 * time.Millisecond)
		}
	}()
	return resultCh
}
