package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genEvenNos()
	for {
		if evenNo, ok := <-ch; ok {
			fmt.Println("even number =", evenNo, ok)
		} else {
			break
		}
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
