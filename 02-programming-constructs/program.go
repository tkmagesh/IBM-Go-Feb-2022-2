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

	/* For construct */
	//version 1.0
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)

	/* version 2.0 (while implementation) */
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Println("numSum =", numSum)

	/* version 3.0 (infinite loop) */
	x := 1
	for {
		x += x
		if x > 100 {
			break
		}
	}
	fmt.Println("x =", x)
}
