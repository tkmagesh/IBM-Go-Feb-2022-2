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
	/*
		products = make([]string, 0, 20)
		fmt.Println(products)
	*/
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

	/* Slicing */
	products = []string{"Pen", "Pencil", "Marker", "Scribble-Pad", "Mouse", "Charger", "iPad", "Adapter"}
	/*
		[lo : hi] => from 'lo' to 'hi-1'
		[:hi] => from 0 to 'hi-1'
		[lo :] => from 'lo' to end

	*/
	fmt.Println("Products =>", products)
	fmt.Println("products[2:5] =>", products[2:5])
	fmt.Println("products[:5] =>", products[:5])
	fmt.Println("products[5:] =>", products[5:])

	productsExceptMarker := append(products[:2], products[3:]...)
	fmt.Println("products except marker =>", productsExceptMarker)

	/* Maps */
	fmt.Printf("\nMaps\n")
	//var productsRank map[string]int = map[string]int{}
	//var productsRank map[string]int = make(map[string]int)

	//productsRank := map[string]int{"Pen": 1, "Pencil": 5, "Marker": 2}
	productsRank := map[string]int{
		"Pen":    1,
		"Pencil": 5,
		"Marker": 2,
	}
	fmt.Printf("%#v\n", productsRank)

	//Adding a new item to a map
	productsRank["Scribble-pad"] = 3
	fmt.Printf("After adding 1 item, productRanks = %#v\n", productsRank)

	//check if an item exists
	keyToCheck := "Phone"
	if val, exists := productsRank[keyToCheck]; exists {
		fmt.Printf("Value of key %q is %v\n", keyToCheck, val)
	} else {
		fmt.Printf("key %q does not exist\n", keyToCheck)
	}

	//remove an item
	delete(productsRank, "Pen")
	fmt.Printf("After deleting \"pen\" productsRank : \n%#v\n", productsRank)

	//Iterate a map
	fmt.Println("Iterating a map")
	for key, val := range productsRank {
		fmt.Printf("Key = %q, Value = %d\n", key, val)
	}
}
