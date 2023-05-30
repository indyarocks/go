package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

func productDetail(product *Product) {
	fmt.Println("Name:", product.name, "Category:", product.category, "Price:", product.price)
}

func (product *Product) methodProductDetail() {
	fmt.Println("Name:", product.name, "Category:", product.category, "Price:", product.price)
}

func main() {
	products := []*Product{
		{"iPhone", "Mobile", 1299},
		{"OnePlus", "Mobile", 699},
		{"Xiaomi", "Mobile", 299},
	}

	for _, product := range products {
		productDetail(product)
		product.methodProductDetail()
	}
}
