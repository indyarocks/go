package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

type SupplierItem struct {
	name, category string
	price          float64
	*Supplier
}

type Supplier struct {
	name, city string
}

func writeName(val struct {
	name, category string
	price          float64
}) {
	fmt.Println("Name:", val.name)
}

func calcTax(product *Product) *Product {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
	return product
}

func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func newSupplierItem(name, category string, price float64, supplier *Supplier) *SupplierItem {
	return &SupplierItem{name, category, price, supplier}
}

func copySupplierItem(si *SupplierItem) SupplierItem {
	si_copy := *si
	si_supplier := *si.Supplier
	si_copy.Supplier = &si_supplier
	return si_copy
}

func main() {
	p1 := Product{name: "Dota", category: "Esports", price: 0}
	p2 := Product{name: "League of Legends", category: "Esports", price: 10}

	array := [1]ProductStock{
		{
			Product:   p1,
			Alternate: p2,
			count:     100,
		},
	}
	fmt.Println("Array:", array[0].Product.name)

	slice := []ProductStock{
		{
			Product:   p1,
			Alternate: p2,
			count:     100,
		},
	}
	fmt.Println("Slice:", slice[0].Product.name)

	kvp := map[string]ProductStock{
		"first": {
			Product:   p1,
			Alternate: p2,
			count:     100,
		},
	}
	fmt.Println("Map:", kvp["first"].Product.name)
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
	p4 := p1
	p1.name = "Updated Name"
	fmt.Println("P1", p1.name)
	fmt.Println("P4", p4.name)
	p5 := &p1
	p1.name = "Dota 2"
	fmt.Println("P1", p1.name)
	fmt.Println("P5", (*p5).name)

	myProduct := calcTax(&Product{
		name:     "iPhone",
		category: "Computers",
		price:    3999,
	})
	fmt.Println("Name", myProduct.name, "Category", myProduct.category, "Price", myProduct.price)
	testProduct := Product{"Test", "test", 100}
	fmt.Println("Name", testProduct.name, "Category", testProduct.category, "Price", testProduct.price)
	products := [2]*Product{
		newProduct("Dota 2", "Esports", 0),
		newProduct("Adidas", "fashion", 100),
	}
	for _, p := range products {
		fmt.Println("Name", p.name, "Category:", p.category, "Price", p.price)
	}

	mySupplier := &Supplier{"Chandan", "Pune"}
	supplierItems := [2]*SupplierItem{
		newSupplierItem("SI1", "C1", 100, mySupplier),
		newSupplierItem("SI2", "C2", 100, mySupplier),
	}
	for _, si := range supplierItems {
		fmt.Println("Name", si.name, "Supplier", si.Supplier.name, si.Supplier.city)
	}

	si1 := newSupplierItem("SI1", "C1", 100, mySupplier)
	si2 := *si1
	si1.Supplier.name = "Test Supplier"
	for _, si := range []SupplierItem{*si1, si2} {
		fmt.Println("Name", si.name, "Supplier", si.Supplier.name, si.Supplier.city)
	}
	si3 := copySupplierItem(si1)
	si1.Supplier.name = "Test Supplier 2"
	for _, si := range []SupplierItem{*si1, si3} {
		fmt.Println("Name", si.name, "Supplier", si.Supplier.name, si.Supplier.city)
	}
}
