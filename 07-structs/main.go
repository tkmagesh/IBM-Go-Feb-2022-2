package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
	Salary    float32
}

func main() {
	//var emp Employee
	//var emp Employee = Employee{}
	//var emp Employee = Employee{100, "Magesh", "Kuppan", 10000}
	//var emp Employee = Employee{Id: 100, FirstName: "Magesh", LastName: "Kuppan"}
	/*
		var emp Employee = Employee{
			Id:        100,
			LastName:  "Kuppan",
			FirstName: "Magesh",
			Salary:    10000,
		}
	*/
	emp := Employee{
		Id:        100,
		LastName:  "Kuppan",
		FirstName: "Magesh",
		Salary:    10000,
	}
	fmt.Printf("%#v\n", emp)

	employees := []Employee{
		{101, "Magesh", "Kuppan", 10000},
		{102, "Suresh", "Kannan", 20000},
		{103, "Rajesh", "Pandit", 30000},
	}
	fmt.Println("Employees")
	for _, emp := range employees {
		fmt.Print(Format(&emp))
	}
}

func Format(emp *Employee) string {
	return fmt.Sprintf("Id=%d, FirstName=%s, LastName=%s, Salary=%f\n", emp.Id, emp.FirstName, emp.LastName, emp.Salary)
}
