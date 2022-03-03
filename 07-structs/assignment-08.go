package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func main() {
	pen := Product{1000, "Pen", 10, 100, "Stationary"}
	fmt.Println(Format(pen))
	ApplyDiscount(&pen, 10)
	fmt.Println("After applying discount:")
	fmt.Println(Format(pen))
}

func Format(p Product) string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%f, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func ApplyDiscount(p *Product, discountPercentage float32) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
