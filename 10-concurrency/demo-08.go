package main

import "fmt"

func main() {
	var ch chan int = make(chan int)
	fmt.Println("main started")
	go func() {
		fmt.Println("Attempting to write the data")
		ch <- 100
	}()
	fmt.Println("Attempting to read the data")
	data := <-ch
	fmt.Println(data)
	fmt.Println("main completed")
}
