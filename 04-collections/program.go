package main

import "fmt"

func main() {
	//Arrays
	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	//var nos = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	//Iterating an array [using indexer]
	fmt.Println("Iterating using an indexer")
	for i := 0; i < len(nos); i++ {
		fmt.Println(nos[i])
	}

	//Iterating an array [using range]
	fmt.Println("Iterating using range")
	for idx, val := range nos {
		fmt.Printf("Value @ [%d] = %d\n", idx, val)
	}

}
