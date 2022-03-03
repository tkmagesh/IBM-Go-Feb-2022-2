package main

import "fmt"

func main() {
	var n1 int = 100
	var n1Ptr *int = &n1 // value to address(pointer)
	fmt.Println(n1, n1Ptr)

	n2 := *n1Ptr //address(pointer) to value - DEREFERENCING
	fmt.Println(n2)

	var x int = 100
	fmt.Println("before incrementing, x = ", x)
	increment(&x)
	fmt.Println("after incrementing, x = ", x)

	a := 100
	b := 200
	fmt.Printf("Before swapping, a = %d and b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swapping, a = %d and b = %d\n", a, b)
}

func increment(n *int) {
	*n = *n + 1
}

func swap(n1, n2 *int) {
	*n1, *n2 = *n2, *n1
}
