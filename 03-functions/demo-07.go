package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("divisor cannot be 0")

func main() {
	q, r, err := divide(100, 0)
	if err == DivideByZeroError {
		fmt.Println("Cannot divide by zero")
		return
	}
	if err != nil {
		fmt.Println("Something went wrong!", err)
		return
	}
	fmt.Printf("Quotient = %d, Remainder = %d\n", q, r)
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := errors.New("divisor cannot be 0")
		return 0, 0, err
	}
	quotient := x / y
	remainder := x % y
	return quotient, remainder, nil

}
*/

func divide(x, y int) (quotient int, remainder int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	quotient = x / y
	remainder = x % y
	return

}
