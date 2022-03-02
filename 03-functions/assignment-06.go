/* panic & recovery */

/*
Modify the program in such a way that when a panic occurs, the program prints a message and continues with the loop
	Important:
		1. DONOT write the logic for checking if y == 0
		2. DONOT modify the "divide" function
*/
package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("cannot divide by zero")

func main() {
	var n1, n2 int
	for {
		fmt.Scanln(&n1, &n2)
		quotient, remainder, err := divideClient(n1, n2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Result = %d, %d\n", quotient, remainder)
	}

}

func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("Divide connot complete")
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
