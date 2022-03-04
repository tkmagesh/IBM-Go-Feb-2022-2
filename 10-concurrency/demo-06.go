package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var opCount int64

func main() {
	wg := &sync.WaitGroup{}
	fmt.Println("main started")
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go f1(i, wg)
	}
	f2()
	wg.Wait()
	fmt.Println("main completed")
	fmt.Println("OpCount = ", opCount)
}

func f1(no int, wg *sync.WaitGroup) {
	atomic.AddInt64(&opCount, 1)
	fmt.Printf("f1[%d] invoked\n", no)
	wg.Done()
}

func f2() {
	fmt.Println("f2 invoked")
}
