/* panic & recovery */
package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("cannot divide by zero")

func main() {
	defer func() {
		fmt.Println("main completed")
		if err := recover(); err != nil {
			fmt.Println("Recovering from panic :", err)
			return
		}
		fmt.Println("Successful execution of the app!")
	}()
	quotient, remainder := divide(100, 0)
	fmt.Printf("Quotient = %d, Remainder = %d\n", quotient, remainder)

}

func divide(x, y int) (quotient, remainder int) {
	if y == 0 {
		panic(DivideByZeroError)
	}
	quotient = x / y
	remainder = x % y
	return
}
