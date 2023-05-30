package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func writeName(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("Name:", val.name)
}

func main() {
	type Product struct {
		name, category string
		price          float64
	}

	type ProductStock struct {
		Product
		Alternate Product
		count     int
	}

	type Item struct {
		name, category string
		price          float64
	}

	p1 := Product{name: "Dota", category: "Esports", price: 0}
	p2 := Product{name: "League of Legends", category: "Esports", price: 10}
	itemStock := ProductStock{
		Product:   p1,
		Alternate: p2,
		count:     100,
	}
	fmt.Println("Product", itemStock.Product.name, "category", itemStock.Product.price, "Count", itemStock.count)
	fmt.Println("Alt Product", itemStock.Alternate.name, "category", itemStock.Alternate.price, "Count", itemStock.count)
	item1 := Item{name: "Dota", category: "Esports", price: 0}

	fmt.Println("p1 == item1", p1 == Product(item1))
	writeName(item1)
	writeName(p1)
	writeName(p2)

	var builder strings.Builder
	json.NewEncoder(&builder).Encode(struct {
		ProductName  string
		ProductPrice float64
	}{
		ProductName:  p1.name,
		ProductPrice: p1.price,
	})
	fmt.Println(builder.String())
}
