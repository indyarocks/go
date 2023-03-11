package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Salary   int
	FullTime bool
}

func main() {
	jack := Employee{
		Name:     "Jack Smith",
		Age:      30,
		Salary:   40000,
		FullTime: false,
	}

	john := Employee{
		Name:     "John Doe",
		Age:      40,
		Salary:   500000,
		FullTime: true,
	}

	var employees []Employee
	employees = append(employees, jack)
	employees = append(employees, john)

	for _, emp := range employees {
		fmt.Println(emp.Name, emp.Salary, emp.Age, emp.FullTime)
	}
}
