package main

import "strconv"

type Product struct {
	Name, Category string
	Price          float64
}

type ProductSlice []*Product

var slice = ProductSlice{
	{"iPhone", "Mobile", 1499},
	{"OnePlus", "Mobile", 799},
	{"AirPod", "Accessories", 499},
	{"Macbook Pro", "Laptop", 2299},
	{"Macbook Air", "Laptop", 1799},
	{"iMac", "Desktop", 2099},
	{"Xiomi", "Mobile", 499},
	{"Dell", "Desktop", 899},
	{"Home", "SelfCare", 0},
}

func ToCurrency(price float64) string {
	return "$" + strconv.FormatFloat(price, 'f', 2, 64)
}
