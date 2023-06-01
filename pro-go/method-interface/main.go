package main

import (
	"fmt"
)

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

type Account struct {
	accountNumber int
	expenses      []Expense
}

func usualPrintDetails(product *Product) {
	fmt.Println("Name:", product.name, "Category:", product.category, "Price:", product.price)
}

func (product *Product) printDetails() {
	fmt.Println("Name:", product.name, "Category:", product.category, "Price:", product.calcTax(500, 0.2))
}

func (product *Product) calcTax(threshold float64, rate float64) float64 {
	if product.price > threshold {
		product.price += product.price * rate
	}
	return product.price
}

func (supplier *Supplier) printDetails() {
	fmt.Println("Supplier Name:", supplier.name, "City:", supplier.city)
}

func getProducts() []Product {
	return []Product{
		{"Airpod", "Accessories", 199},
		{"iPhone", "Mobile", 1499},
		{"iPad", "Mobile", 1199},
		{"Apple Watch", "Watch", 599},
	}
}

func (products *ProductList) calcCategoryTotal() map[string]float64 {
	total := make(map[string]float64)
	for _, product := range *products {
		total[product.category] += product.price
	}
	return total
}

func calcTotal(expenses []Expense) (total float64) {
	for _, expense := range expenses {
		total += expense.getCost(true)
	}
	return
}

func processItem(item interface{}) {
	switch value := item.(type) {
	case Service:
		fmt.Println("Service:", value.description, "Price:", value.monthlyFee*(float64(value.durationMonth)))
	case Product:
		fmt.Println("Product:", value.name, "Price:", value.price)
	case *Service:
		fmt.Println("Service:", value.description, "Price:", value.monthlyFee*(float64(value.durationMonth)))
	case *Product:
		fmt.Println("Product:", value.name, "Price:", value.price)
	case string, bool, int:
		fmt.Println("Built in type", value)
	default:
		fmt.Println("Default", value)
	}
}

func main() {
	products := []*Product{
		{"iPhone", "Mobile", 1299},
		{"OnePlus", "Mobile", 699},
		{"Xiaomi", "Mobile", 299},
	}

	for _, product := range products {
		usualPrintDetails(product)
		product.printDetails()
	}

	myProduct := Product{"Ruby Hero", "Programming", 99999999}
	myProduct.printDetails()

	suppliers := []*Supplier{
		{"Ali Baba", "China"},
		{"Amazon", "US"},
		{"flipkart", "India"},
	}

	for _, supplier := range suppliers {
		supplier.printDetails()
	}
	(*Supplier).printDetails(&Supplier{"Ali Baba", "China"})
	(*Product).printDetails(&Product{"iPhone", "Mobile", 1299})

	productList := ProductList(getProducts())
	for category, total := range productList.calcCategoryTotal() {
		fmt.Println("For category", category, "Total", total)
	}

	insurance := Service{"LIC", 12, 1200}
	fmt.Println("Service", insurance.description, "Yearly Fee", (insurance.monthlyFee)*float64(insurance.durationMonth))

	expenses := []Expense{
		Product{"iPhone", "Mobile", 1499},
		Service{"Insurance", 12, 1200},
	}

	for _, expense := range expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}
	fmt.Println("Total Expense", calcTotal(expenses))

	account := Account{
		accountNumber: 101,
		expenses: []Expense{
			Product{"Kindle", "Tablet", 299},
			Service{"Painting", 1, 500},
		},
	}

	for _, expense := range account.expenses {
		fmt.Println("Expense", expense.getName(), "Cost:", expense.getCost(true))
	}
	fmt.Println("Total", calcTotal(account.expenses))

	serviceExpenses := []Expense{
		Service{"Insurance", 12, 1200},
		Service{"Medical", 12, 200},
		Service{"Luxury", 12, 5000},
		//Product{"Test Assertion", "Debug", 100}, // It'll error out in panic
	}
	for _, expense := range serviceExpenses {
		s := expense.(Service) // Dynamic Type Assertion
		fmt.Println("Service:", s.description, "Price:", s.monthlyFee*(float64(s.durationMonth)))
	}

	expensesWithTypeAssertionCheck := []Expense{
		Service{"Insurance", 12, 1200},
		Service{"Medical", 12, 200},
		&Service{"Luxury", 12, 5000},
		&Product{"Test Assertion", "Debug", 100},
		Product{"Test Assertion1", "Debug1", 1000},
	}

	for _, expense := range expensesWithTypeAssertionCheck {
		if s, ok := expense.(Service); ok {
			fmt.Println("Service:", s.description, "Price:", s.monthlyFee*(float64(s.durationMonth)))
		} else {
			fmt.Println("Expense", expense.getName(), "Cost", expense.getCost(true))
		}
	}

	expensesWithTypeAssertionCheckSwitch := []Expense{
		Service{"Insurance", 12, 1200},
		Service{"Medical", 12, 200},
		&Service{"Luxury", 12, 5000},
		&Product{"Test Assertion", "Debug", 100},
		Product{"Test Assertion1", "Debug1", 1000},
	}

	for _, expense := range expensesWithTypeAssertionCheckSwitch {
		switch value := expense.(type) {
		case Service:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee*(float64(value.durationMonth)))
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		case *Service:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee*(float64(value.durationMonth)))
		case *Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		default:
			fmt.Println("Expense", expense.getName(), "Cost", expense.getCost(true))
		}
	}

	fmt.Println("Empty interface==================")
	data := []interface{}{
		insurance,
		Service{"Insurance", 12, 1200},
		Service{"Medical", 12, 200},
		&Service{"Luxury", 12, 5000},
		&Product{"Test Assertion", "Debug", 100},
		Product{"Test Assertion1", "Debug1", 1000},
		expenses,
		"Lol string",
		100,
		true,
	}

	for _, item := range data {
		processItem(item)
	}
}
