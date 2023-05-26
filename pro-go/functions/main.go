package main

import "fmt"

func productSuppliersSlice(product string, suppliers []string) {
	for _, supplier := range suppliers {
		fmt.Println("Product", product, "supplier", supplier)
	}
}

func productSuppliersVariadic(product string, suppliers ...string) {
	if len(suppliers) == 0 {
		fmt.Println("Product", product, "supplier", "(None)")
	} else {
		for _, supplier := range suppliers {
			fmt.Println("Product", product, "supplier", supplier)
		}
	}
}

func swapInts(x, y int) {
	fmt.Println("Inside function:", "x:", x, "y:", y)
	var temp int = x
	x = y
	y = temp
	fmt.Println("Inside function:", "x:", x, "y:", y)
}

func swapIntPointers(x, y *int) {
	fmt.Println("Inside pointer function:", "x:", *x, "y:", *y)
	var temp int = *x
	*x = *y
	*y = temp
	fmt.Println("Inside pointer function:", "x:", *x, "y:", *y)
}

func calTotal(products map[string]float64) (count int, total float64) {
	fmt.Println("Function calTotal Started")
	defer fmt.Println("First defer")
	for _, price := range products {
		count++
		total += price
	}
	defer fmt.Println("Second defer")
	fmt.Println("Function returning")
	return
}

func main() {
	fmt.Println("Slice argument...")
	productSuppliersSlice("iPhone", []string{"Apple", "Axios"})
	productSuppliersSlice("iPhone", []string{"Apple Inc"})
	fmt.Println("Variadic function...")
	productSuppliersVariadic("iPhone", "Apple", "Axios")
	productSuppliersVariadic("iPhone", "Apple Inc")
	productSuppliersVariadic("iPhone")
	suppliers := []string{"Apple", "Microsoft", "IBM"}
	productSuppliersVariadic("iPhone", suppliers...)
	var x, y int = 1, 100
	swapInts(x, y)
	fmt.Println("outside function:", "x:", x, "y:", y)
	swapIntPointers(&x, &y)
	fmt.Println("outside pointer function:", "x:", x, "y:", y)

	products := map[string]float64{
		"iPhone14":       1299.00,
		"iPhone14Pro":    1599.00,
		"iPhone14ProMax": 1899,
	}
	prodCount, total := calTotal(products)
	fmt.Println("Total Count:", prodCount, "total amount", total)
}
