package main

import "fmt"

func calcWithTax(price float64) float64 {
	return price * 0.1
}

func calcWithoutTax(price float64) float64 {
	return 0
}

func printPrice(product string, price float64, taxCalculator func(float642 float64) float64) {
	finalPrice := price + taxCalculator(price)
	fmt.Println("====Product", product, "Final Price", finalPrice)
}

func selectCalculator(price float64) func(float64) float64 {
	if price > 300 {
		return calcWithTax
	} else {
		return calcWithoutTax
	}
}

func main() {
	products := map[string]float64{
		"iPhone":  1499,
		"OnePlus": 699,
		"Mi":      299,
	}
	for product, price := range products {
		var calTax func(float64) float64
		var calTaxUsingSelector func(float64) float64 = selectCalculator(price)

		fmt.Println("function calTax assigned", calTax != nil)
		if price > 300 {
			calTax = calcWithTax
		} else {
			calTax = calcWithoutTax
		}

		fmt.Println("function calTax assigned", calTax != nil)
		tax := calTax(price)
		fmt.Println("product", product, "Tax amount", tax)
		printPrice(product, price, calTax)
		printPrice(product, price, calTaxUsingSelector)
	}

}
