package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) Format() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%f, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func (p *Product) ApplyDiscount(discountPercentage float32) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

/* type PerishableProduct struct {
	Product //composition
	Expiry  string
} */

func main() {

	pen := Product{100, "Pen", 10, 100, "Stationary"}
	/*
		fmt.Println(Format(pen))
		ApplyDiscount(&pen, 10)
	*/

	fmt.Println(pen.Format())
	pen.ApplyDiscount(10)
	fmt.Println(pen.Format())

	pencilPtr := &Product{200, "Pencil", 1, 500, "Stationary"}
	fmt.Println(pencilPtr.Format())
	pencilPtr.ApplyDiscount(10)
	fmt.Println(pencilPtr.Format())

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
}
