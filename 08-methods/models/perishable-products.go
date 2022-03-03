package models

import "fmt"

type PerishableProduct struct {
	Product //composition
	Expiry  string
}

//overriding the Format method
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry=%s", pp.Product.Format(), pp.Expiry)
}

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:       id,
			Name:     name,
			Cost:     cost,
			Units:    units,
			Category: category,
		},
		Expiry: expiry,
	}
}
