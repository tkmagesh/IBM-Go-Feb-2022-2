package main

import (
	"fmt"
	"time"
)

//consumer
func main() {
	done := make(chan bool)
	ch := genEvenNos(done)
	go func() {
		var input string
		fmt.Println("Hit ENTER to stop...")
		fmt.Scanln(&input)
		done <- true
	}()
	for evenNo := range ch {
		fmt.Println("even number =", evenNo)
	}
}

//producer
func genEvenNos(done chan bool) chan int {
	resultCh := make(chan int)
	go func() {
		counter := 1
	LOOP:
		for {
			select {
			case resultCh <- counter * 2:
				counter++
				time.Sleep(500 * time.Millisecond)
			case <-done:
				break LOOP
			}
		}
		close(resultCh)
	}()
	return resultCh
}
