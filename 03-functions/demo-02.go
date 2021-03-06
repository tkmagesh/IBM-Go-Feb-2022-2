/* variadic functions */
package main

import "fmt"

func main() {
	fmt.Println(sum(10))
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, 20, 30))
	fmt.Println(sum(10, 20, 30, 40, 50))
}

func sum(nos ...int) (result int) {
	for i := 0; i < len(nos); i++ {
		result += nos[i]
	}
	return
}
