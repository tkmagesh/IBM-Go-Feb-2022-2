package main

import "fmt"

func main() {
	/* if else statement */

	/*
		var no = 21
		if no%2 == 0 {
			fmt.Println(no, "is Even")
		} else {
			fmt.Println(no, "is Odd")
		}
		fmt.Println("No is ", no)
	*/

	if no := 21; no%2 == 0 {
		fmt.Println(no, "is Even")
	} else {
		fmt.Println(no, "is Odd")
	}
	//fmt.Println("No is ", no) // will not work as the scope of the "no" variable is limited to "if else" block
}
