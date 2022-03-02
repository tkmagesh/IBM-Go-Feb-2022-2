/* anonymous functions */

package main

import "fmt"

func main() {
	func() {
		fmt.Println("fn invoked")
	}()

	func(x, y int) {
		fmt.Printf("Adding %d and %d, result = %d\n", x, y, x+y)
	}(100, 200)

	result := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Printf("result = %d\n", result)
}
