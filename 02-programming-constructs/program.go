package main

import "fmt"

func main() {
	/* if else statement */

	/*
		var no = 21
		if no%2 == 0 {
			fmt.Println(no, "is Even")
		} else {
			fmt.Println(no, "is Odd")
		}
		fmt.Println("No is ", no)
	*/

	if no := 21; no%2 == 0 {
		fmt.Println(no, "is Even")
	} else {
		fmt.Println(no, "is Odd")
	}
	//fmt.Println("No is ", no) // will not work as the scope of the "no" variable is limited to "if else" block

	/* For construct */
	//version 1.0
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)

	/* version 2.0 (while implementation) */
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Println("numSum =", numSum)

	/* version 3.0 (infinite loop) */
	x := 1
	for {
		x += x
		if x > 100 {
			break
		}
	}
	fmt.Println("x =", x)

	/* switch case construct */
	/*
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Very good"
		score 10 => "Excellent"
		otherwise => "Invalid score"
	*/

	score := 6
	/*
		switch score {
		case 0:
			fmt.Println("Terrible")
		case 1:
			fmt.Println("Terrible")
		case 2:
			fmt.Println("Terrible")
		case 3:
			fmt.Println("Terrible")
		case 4:
			fmt.Println("Not Bad")
		case 5:
			fmt.Println("Not Bad")
		case 6:
			fmt.Println("Not Bad")
		case 7:
			fmt.Println("Not Bad")
		case 8:
			fmt.Println("Very Good")
		case 9:
			fmt.Println("Very Good")
		case 10:
			fmt.Println("Excellent")
		default:
			fmt.Println("Invalid score")

		}
	*/

	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not Bad")
	case 8, 9:
		fmt.Println("Very Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid score")
	}

	num := 21
	switch {
	case num%2 == 0:
		fmt.Println(num, "is even")
	case num%2 == 1:
		fmt.Println(num, "is odd")
	}

	//falthrough
	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
	case 1:
		fmt.Println("n <= 1")
		fallthrough
	case 2:
		fmt.Println("n <= 2")
		fallthrough
	case 3:
		fmt.Println("n <= 3")
		fallthrough
	case 4:
		fmt.Println("n <= 4")
		fallthrough
	case 5:
		fmt.Println("n <= 5")
		fallthrough
	case 6:
		fmt.Println("n <= 6")
	case 7:
		fmt.Println("n <= 7")
		fallthrough
	case 8:
		fmt.Println("n <= 8")
	}
}
