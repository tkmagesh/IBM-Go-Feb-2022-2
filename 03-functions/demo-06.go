/*
	Higher Order functions
		3. Returning a function as result from other functions
*/
package main

import "fmt"

func main() {
	//Use Case - 1
	adder := getAdder()
	fmt.Println(adder(100, 200)) //=> 300

	//Use Case - 2
	/*
		add(100, 200)
		subtract(100, 200)

		logOperation(100, 200, add)
		logOperation(100, 200, subtract)
	*/

	/* use case - 3 */
	logAdd := getLogOperation(add)
	logAdd(100, 200)

	logSubtract := getLogOperation(subtract)
	logSubtract(100, 200)
}

func getAdder() func(int, int) int {
	add := func(x, y int) int {
		return x + y
	}
	return add
}

func getLogOperation(oper func(int, int)) func(int, int) {
	return func(x, y int) {
		fmt.Println("Before invocation")
		oper(x, y)
		fmt.Println("After invocation")
	}
}

func add(x, y int) {
	fmt.Println("Add Result = ", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result = ", x-y)
}
