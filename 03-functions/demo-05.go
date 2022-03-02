/*
Higher Order functions
	2. pass functions as arguments to other functions
*/
package main

import "fmt"

func main() {

	//step-1
	/*
		fmt.Println("Before invocation")
		add(100, 200)
		fmt.Println("After invocation")

		fmt.Println("Before invocation")
		subtract(100, 200)
		fmt.Println("After invocation")
	*/
	x, y := 100, 200
	//step-2
	/*
		logAdd(x, y)
		logSubtract(x, y)
	*/

	//step-3
	logOperation(x, y, add)
	logOperation(x, y, subtract)

	logOperationV2(x, y, multiply)
	logOperationV2(x, y, divide)
}

func logOperation(x, y int, oper func(int, int)) {
	fmt.Println("Before invocation")
	oper(x, y)
	fmt.Println("After invocation")
}

func logOperationV2(x, y int, operV2 func(int, int) int) {
	fmt.Println("Before invocation")
	result := operV2(x, y)
	fmt.Println("Result = ", result)
	fmt.Println("After invocation")
}

/*
func logAdd(x, y int) {
	fmt.Println("Before invocation")
	add(x, y)
	fmt.Println("After invocation")
}

func logSubtract(x, y int) {
	fmt.Println("Before invocation")
	subtract(x, y)
	fmt.Println("After invocation")
}
*/

func add(x, y int) {
	fmt.Println("Add Result = ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result = ", x-y)
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
