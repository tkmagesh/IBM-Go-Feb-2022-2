package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genEvenNos()
	for evenNo := range ch {
		fmt.Println("even number =", evenNo)
	}
}

func genEvenNos() chan int {
	resultCh := make(chan int)
	go func() {
		for i := 1; i <= 20; i++ {
			evenNo := i * 2
			resultCh <- evenNo
			time.Sleep(500 * time.Millisecond)
		}
		close(resultCh)
	}()
	return resultCh
}
