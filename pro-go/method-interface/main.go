package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

type Supplier struct {
	name, city string
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

	suppliers := []*Supplier{
		{"Ali Baba", "China"},
		{"Amazon", "US"},
		{"flipkart", "India"},
	}

	for _, supplier := range suppliers {
		supplier.printDetails()
	}
}
