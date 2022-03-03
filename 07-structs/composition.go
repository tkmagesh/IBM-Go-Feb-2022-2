package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

type PerishableProduct struct {
	Product //composition
	Expiry  string
}

func main() {
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
}

func Format(p Product) string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%f, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func ApplyDiscount(p *Product, discountPercentage float32) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
