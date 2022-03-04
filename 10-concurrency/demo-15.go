package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		evenCh := genEvenNos()
		for evenNo := range evenCh {
			fmt.Println("even number =", evenNo)
		}
		wg.Done()
	}()

	go func() {
		oddCh := genOddNos()
		for oddNo := range oddCh {
			fmt.Println("odd number =", oddNo)
		}
		wg.Done()
	}()

	wg.Wait()
}

func genEvenNos() chan int {
	evenCh := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			evenNo := i * 2
			evenCh <- evenNo
			time.Sleep(500 * time.Millisecond)
		}
		close(evenCh)
	}()
	return evenCh
}

func genOddNos() chan int {
	oddCh := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			oddNo := (i * 2) + 1
			oddCh <- oddNo
			time.Sleep(1 * time.Second)
		}
		close(oddCh)
	}()
	return oddCh
}
