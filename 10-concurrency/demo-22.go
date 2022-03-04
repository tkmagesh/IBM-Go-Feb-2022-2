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
	ch1 := make(chan string)
	ch2 := make(chan string)
	go print("Hello", ch1, ch2)
	go print("World", ch2, ch1)
	ch1 <- "start"
	var input string
	fmt.Scanln(&input)
}

func print(str string, x, y chan string) {
	for i := 0; i < 5; i++ {
		<-x
		fmt.Println(str)
		time.Sleep(500 * time.Millisecond)
		y <- "done"
	}
}
