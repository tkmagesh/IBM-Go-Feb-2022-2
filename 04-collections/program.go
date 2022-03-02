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

	fmt.Printf("\nSlice\n")
	//var products []string
	//var products []string = []string{"Pen", "Pencil", "Marker"}
	//var products = []string{"Pen", "Pencil", "Marker"}
	//products := []string{"Pen", "Pencil", "Marker"}

	var products []string
	//products[0] = "Pen"
	//products = append(products, "Scribble-Pad")

	fmt.Printf("\nMemory Allocation\n")

	/* preallocation of memory */
	products = make([]string, 0, 20)
	fmt.Println(products)
	/*
		products = append(products, "Pen")
		fmt.Printf("%v, len=%d, capacity=%d\n", products, len(products), cap(products))
		products = append(products, "Pencil")
		fmt.Printf("%v, len=%d, capacity=%d\n", products, len(products), cap(products))
		products = append(products, "Marker")
		fmt.Printf("%v, len=%d, capacity=%d\n", products, len(products), cap(products))
		products = append(products, "Scribble-Pad")
		fmt.Printf("%v, len=%d, capacity=%d\n", products, len(products), cap(products))
		products = append(products, "Mouse")
		fmt.Printf("%v, len=%d, capacity=%d\n", products, len(products), cap(products))
		fmt.Println(products)
	*/

	data := make([]int, 5, 20)
	fmt.Println(data, len(data), cap(data))

}
