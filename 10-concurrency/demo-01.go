package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")
	go f1()
	f2()
	time.Sleep(1 * time.Millisecond)
	fmt.Println("main completed")
}

func f1() {
	time.Sleep(1 * time.Second)
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
