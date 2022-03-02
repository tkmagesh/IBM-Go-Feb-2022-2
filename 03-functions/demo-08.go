/* deferred functions */

package main

import "fmt"

func main() {
	/*
		defer func() {
			fmt.Println("[main] - deferred")
		}()
	*/
	defer fmt.Println("[main] - deferred")
	fmt.Println("main started")
	result := f1()
	fmt.Println("f1 Result =", result)
	fmt.Println("main completed")
}

func f1() (result int) {
	defer func() {
		fmt.Println("[f1] - deferred - 1")
		result = 1000
	}()
	defer func() {
		fmt.Println("[f1] - deferred - 2")
		result = 2000
	}()
	defer func() {
		fmt.Println("[f1] - deferred - 3")
		result = 3000
	}()
	fmt.Println("f1 started")
	f2()
	fmt.Println("f1 completed")
	return 100
}

func f2() {
	defer deferredF2()
	fmt.Println("f2 started")
	fmt.Println("f2 completed")
}

func deferredF2() {
	fmt.Println("[f2] - deferred")
}
