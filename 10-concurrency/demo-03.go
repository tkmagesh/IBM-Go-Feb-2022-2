package main

import (
	"fmt"
	"sync"
	"time"
)

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
}

func f1(no int, wg *sync.WaitGroup) {
	time.Sleep(5 * time.Second)
	fmt.Printf("f1[%d] invoked\n", no)
	wg.Done()
}

func f2() {
	fmt.Println("f2 invoked")
}
