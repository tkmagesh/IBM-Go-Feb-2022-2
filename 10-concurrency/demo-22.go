package main

import (
	"fmt"
	"time"
)

/* Change the program to behave as follows */
/*
Hello
World
Hello
World
Hello
World
Hello
World
Hello
World

NOTE:
	The loop should be only in the "print" function
	the print function should print only one string
*/

func main() {
	print("Hello")
	print("World")
}

func print(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(str)
		time.Sleep(500 * time.Millisecond)
	}
}
