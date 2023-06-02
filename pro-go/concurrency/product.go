package main

import "strconv"

type Product struct {
	Name, Category string
	Price          float64
}

var ProductList = []*Product{
	{"Kayak", "WaterSports", 275},
	{"LifeJackets", "WaterSports", 500},
	{"Soccer Ball", "Soccer", 20},
	{"Corner Flag", "Soccer", 50},
	{"Stadium", "Soccer", 5000000},
	{"Thinking Cap", "Chess", 50.45},
	{"Unsteady Chair", "Chess", 350.50},
	{"Bling Bling King", "Chess", 1200},
}

type ProductGroup []*Product

type ProductData map[string]ProductGroup

var Products = make(ProductData)

func ToCurrency(val float64) string {
	return "$" + strconv.FormatFloat(val, 'f', 2, 64)
}

func init() {
	for _, product := range ProductList {
		if _, ok := Products[product.Category]; ok {
			Products[product.Category] = append(Products[product.Category], product)
		} else {
			Products[product.Category] = ProductGroup{product}
		}
	}
}
