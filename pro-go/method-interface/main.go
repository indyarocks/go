package main

import "fmt"

type Product struct {
	name, category string
	price          float64
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
}
