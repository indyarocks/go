package main

import (
	"fmt"
	"log"
	"myapp/staff"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Doe", Salary: 30000, FullTime: true},
	{FirstName: "Sally", LastName: "Jones", Salary: 50000, FullTime: true},
	{FirstName: "Mark", LastName: "Smithers", Salary: 60000, FullTime: true},
	{FirstName: "Allan", LastName: "Anderson", Salary: 15000, FullTime: false},
	{FirstName: "Margaret", LastName: "Carter", Salary: 100000, FullTime: false},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}

	log.Println(myStaff.All())
	fmt.Println()
	log.Println(myStaff.OverPaid())
	fmt.Println()
	staff.OverPaidLimit = 10000
	log.Println("OverPaid Staffs", myStaff.OverPaid())
	log.Println("UnderPaid Staffs", myStaff.UnderPaid())

	myMap := make(map[int]string)
	myMap[1] = "hi"
	myMap[2] = "bye"
	log.Println(myMap)
}
