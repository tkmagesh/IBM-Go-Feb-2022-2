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
	var quotient, remainder int
	for {
		fmt.Scanln(&n1, &n2)
		quotient, remainder = divide(n1, n2)
		fmt.Printf("Result = %d, %d\n", quotient, remainder)
	}

}

func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
