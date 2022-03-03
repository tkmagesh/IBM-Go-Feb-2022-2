package main

import (
	"fmt"
	"methods-demo/models"
)

type MyStr string //alias for the string

func (s MyStr) Length() int {
	return len(s)
}

func main() {

	pen := models.Product{
		Id:       100,
		Name:     "Pen",
		Cost:     10,
		Units:    100,
		Category: "Stationary",
	}
	/*
		fmt.Println(Format(pen))
		ApplyDiscount(&pen, 10)
	*/

	fmt.Println(pen.Format())
	pen.ApplyDiscount(10)
	fmt.Println(pen.Format())

	pencilPtr := &models.Product{
		Id:       200,
		Name:     "Pencil",
		Cost:     1,
		Units:    500,
		Category: "Stationary",
	}

	fmt.Println(pencilPtr.Format())
	pencilPtr.ApplyDiscount(10)
	fmt.Println(pencilPtr.Format())
	fmt.Println(pencilPtr.WhoAmI())
	/* compostion */
	/* grapes := models.PerishableProduct{
		models.Product{200, "Grapes", 50, 25, "Fruits"},
		"2 Days",
	} */
	grapes := models.NewPerishableProduct(200, "Grapes", 50, 20, "Fruits", "2 Days")

	fmt.Println(grapes.Format())
	grapes.ApplyDiscount(10)
	fmt.Println(grapes.Format())

	/*
		grapes := PerishableProduct{
			Product{200, "Grapes", 50, 25, "Fruits"},
			"2 Days",
		}
		fmt.Printf("%#v\n", grapes)
		fmt.Println(grapes.Product.Name)
		fmt.Println(grapes.Name)

		fmt.Println(Format(grapes.Product))
		ApplyDiscount(&(grapes.Product), 10)
		fmt.Println("After applying 10% discount")
		fmt.Println(Format(grapes.Product))
	*/

	var s MyStr = MyStr("This is a sample string")
	fmt.Println(s.Length())
}
