package main

import (
	"fmt"
	"time"
)

//consumer
func main() {
	done := time.After(10 * time.Second)
	ch := genEvenNos(done)
	for evenNo := range ch {
		fmt.Println("even number =", evenNo)
	}
}

//producer
func genEvenNos(done <-chan time.Time) <-chan int /* receive only channel */ {
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
