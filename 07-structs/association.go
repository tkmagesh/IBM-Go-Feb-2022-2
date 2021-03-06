package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
	Salary    float32
	Org       Organization
}

type Organization struct {
	Name string
	City string
}

func main() {
	emp := Employee{
		Id:        100,
		FirstName: "Magesh",
		LastName:  "Kuppan",
		Salary:    10000,
		Org: Organization{
			Name: "IBM",
			City: "Bengaluru",
		},
	}
	fmt.Printf("%#v\n", emp)
}
