package main

type Product struct {
	name, category string
	price          float64
}

type ProductList []Product

type Supplier struct {
	name, city string
}

func (product Product) getName() string {
	return product.name
}

func (product Product) getCost(_ bool) float64 {
	return product.price
}
