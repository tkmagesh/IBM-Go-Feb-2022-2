/* Problem statement:
Modify the program to accommodate the following use cases:
	1. count even numbers
	2. count odd numbers
	3. count prime numbers
*/
package main

import "fmt"

func main() {
	fmt.Println(countPrimes(3, 4, 5, 6, 7, 8, 9, 12, 11, 17, 21))
}

func countPrimes(nos ...int) (count int) {
	for i := 0; i < len(nos); i++ {
		if isPrime(nos[i]) {
			count++
		}
	}
	return
}

func isPrime(no int) bool {
	for i := 2; i < no; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
