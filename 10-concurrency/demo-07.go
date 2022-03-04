package main

import "fmt"

//Communicate By Sharing Memory [ NOT ADVICABLE ]
var result int

func main() {
	fmt.Println("main started")
	go add(100, 200)
	fmt.Println("result = ", result)
	fmt.Println("main completed")
}

func add(x, y int) {
	result = x + y
}
