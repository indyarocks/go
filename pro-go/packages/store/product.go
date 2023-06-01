// Package store provides types and methods
// commonly required for online sales
package store

import "fmt"

var standardTax = newTaxRate(0.25, 20)

// Product describes an item for sale
type Product struct {
	Name, Category string
	price          float64
}

func init() {
	fmt.Println("product.go init function invoked")
}

// NewProduct - Product constructor
func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

// Price - Get product Price
func (p *Product) Price() float64 {
	return standardTax.calcPriceWithTax(p)
}

// SetPrice - Set product price
func (p *Product) SetPrice(newPrice float64) {
	p.price = newPrice
}
