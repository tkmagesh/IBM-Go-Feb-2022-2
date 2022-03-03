package main

import (
	"fmt"
	/* "modularity-demo/calculator" */

	calc "modularity-demo/calculator"
	"modularity-demo/calculator/utils"

	"github.com/fatih/color"
)

func main() {
	/*
		fmt.Println(calculator.Add(100, 200))
		fmt.Println(calculator.Subtract(100, 200))
		fmt.Println("OpCount = ", calculator.OpCount())
	*/
	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	color.Green(fmt.Sprintf("OpCount = %d\n", calc.OpCount()))
	fmt.Println("Is 201 even ? ", utils.IsEven(201))
	fmt.Println("Is 201 odd ? ", utils.IsOdd(201))
	color.Red("This line will be printed in red color")
}
