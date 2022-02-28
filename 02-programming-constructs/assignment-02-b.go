package main

import (
	"fmt"
)

func main() {
	var userChoice int
	var n1, n2, result int
LOOP:
	for {
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Println("Enter your choice:")
		fmt.Scanln(&userChoice)

		switch userChoice {
		case 1:
			fmt.Println("Enter the operands (seperated by space):")
			fmt.Scanf("%d %d\n", &n1, &n2)
			result = n1 + n2
			fmt.Printf("Result = %d\n", result)
		case 2:
			fmt.Println("Enter the operands (seperated by space):")
			fmt.Scanf("%d %d\n", &n1, &n2)
			result = n1 - n2
			fmt.Printf("Result = %d\n", result)
		case 3:
			fmt.Println("Enter the operands (seperated by space):")
			fmt.Scanf("%d %d\n", &n1, &n2)
			result = n1 * n2
			fmt.Printf("Result = %d\n", result)
		case 4:
			fmt.Println("Enter the operands (seperated by space):")
			fmt.Scanf("%d %d\n", &n1, &n2)
			result = n1 / n2
			fmt.Printf("Result = %d\n", result)
		case 5:
			break LOOP
		default:
			fmt.Println("Invalid choice")
		}

	}
}
