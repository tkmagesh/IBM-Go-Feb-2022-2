package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("main started")
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go f1(i)
	}
	f2()
	wg.Wait()

	fmt.Println("main completed")
}

func f1(no int) {
	time.Sleep(5 * time.Second)
	fmt.Printf("f1[%d] invoked\n", no)
	wg.Done()
}

func f2() {
	fmt.Println("f2 invoked")
}
