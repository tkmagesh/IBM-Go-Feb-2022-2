package main

import "fmt"

/* Variables declared at the package level are accessible across functions within the same package */

var msg string = "Hello world from package"

func main() {
	/*
		var msg string
		msg = "Hello World!"
	*/

	/*
		var msg string = "Hello World!"
	*/

	//Type Inference
	/*
		var msg = "Hello World!"
	*/

	//msg := "Hello World!"   // := can be used ONLY in a function (not at package level)

	fmt.Println(msg)
	fn()
	fmt.Println(msg)

	/* Declaring multiple variables */
	/* Version 1.0 */
	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Result of addition = "
		result = x + y
		fmt.Println(x, y, str, result)
	*/

	/* Version 2.0 */
	/*
		var x int = 100
		var y int = 200
		var str string = "Result of addition = "
		var result int = x + y
		fmt.Println(x, y, str, result)
	*/

	/* Version 3.0 */
	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Result of addition = "
		result = x + y
		fmt.Println(x, y, str, result)
	*/

	/* Version 4.0 */
	/*
		var x, y int = 100, 200
		var result int = x + y
		var str string = "Result of addition = "
		fmt.Println(x, y, str, result)
	*/

	/* Version 5.0 */
	/*
		var (
			x      int
			y      int
			str    string
			result int
		)
		x = 100
		y = 200
		str = "Result of addition = "
		result = x + y
		fmt.Println(x, y, str, result)
	*/

	/* Version 6.0 */
	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Result of addition = "
		result = x + y
		fmt.Println(x, y, str, result)
	*/

	/* Version 7.0 */
	/*
		var (
			x, y   int    = 100, 200
			result int    = x + y
			str    string = "Result of addition = "
		)
		fmt.Println(x, y, str, result)
	*/

	/* Version 8.0 */
	/*
		var (
			x, y   = 100, 200
			result = x + y
			str    = "Result of addition = "
		)
		fmt.Println(x, y, str, result)
	*/

	/* Version 9.0 */
	/*
		var (
			x, y, str = 100, 200, "Result of addition = "
			result    = x + y
		)
		fmt.Println(x, y, str, result)
	*/

	/* Version 10.0 */
	/*
		x, y := 100, 200
		result := x + y
		str := "Result of addition = "
		fmt.Println(x, y, str, result)
	*/

	/* Version 11.0 */

	x, y, str := 100, 200, "Result of addition = "
	result := x + y
	fmt.Println(x, y, str, result)

}

func fn() {
	msg = "Hello world from fn"
}
