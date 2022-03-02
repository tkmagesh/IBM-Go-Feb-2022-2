/* functions basic */

package main

import "fmt"

func main() {
	fn()
	greet("Magesh")
	fmt.Printf("Is 27 a prime number ? %t\n", isPrime(27))
	fmt.Printf("Is 97 a prime number ? %t\n", isPrime(97))
	fmt.Printf("Add result of 100 and 200 = %d\n", add(100, 200))
	fmt.Printf("Add(2) result of 100 and 200 = %d\n", add2(100, 200))
	/*
		fmt.Println(divide(100, 7))
		fmt.Println(divide2(100, 7))
	*/

	/*
		quotient, remainder := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", quotient, remainder)
	*/

	quotient, _ := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d\n", quotient)
}

/* No arguments, No return values */
func fn() {
	fmt.Println("fn invoked")
}

/* 1 argument, No return values */
func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

/* 1 argument, 1 return value */
func isPrime(no int) bool {
	for i := 2; i < no; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

/* 2 arguments, 1 return value */
func add(x int, y int) int {
	return x + y
}

/* 2 arguments, 1 named return value */
func add2(x, y int) (result int) {
	result = x + y
	return
}

/* 2 arguments, 2 return values */
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}

/* 2 arguments, 2 named return values */
func divide2(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
