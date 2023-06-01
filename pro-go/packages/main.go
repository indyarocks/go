package main

import (
	"fmt"
	"github.com/fatih/color"
	_ "packages/data"
	currencyFmt "packages/fmt"
	"packages/store"
	"packages/store/cart"
)

func main() {
	product := store.NewProduct("iPhone", "Mobile", 1299)

	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", currencyFmt.ToCurrency(product.Price()))

	cart := cart.Cart{
		CustomerName: "Chandan",
		Product: []store.Product{
			*product,
		},
	}
	fmt.Println("Customer Name:", cart.CustomerName)
	fmt.Println("Total:", currencyFmt.ToCurrency(cart.GetTotal()))
	color.Green("Customer Name:" + cart.CustomerName)
	color.Red("Total:" + currencyFmt.ToCurrency(cart.GetTotal()))
}
