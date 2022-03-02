/*
	Higher order functions
	1. functions can be assigned to variables
*/
package main

import "fmt"

func main() {
	/*
		fn := func() {
			fmt.Println("fn invoked")
		}
	*/
	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 200))
}
