package main

import (
	"fmt"
	"sync"
)

type OpCount struct {
	count int
	sync.Mutex
}

func (op *OpCount) Increment() {
	op.Lock()
	{
		op.count++
	}
	op.Unlock()
}

var opCount OpCount

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
	fmt.Println("OpCount = ", opCount.count)
}

func f1(no int, wg *sync.WaitGroup) {
	opCount.Increment()
	fmt.Printf("f1[%d] invoked\n", no)
	wg.Done()
}

func f2() {
	fmt.Println("f2 invoked")
}
