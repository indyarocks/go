package main

import "fmt"

func CalcStoreTotal(data ProductData) float64 {
	var storeTotal float64
	for category, productGroup := range data {
		storeTotal += productGroup.GetCategoryPrice(category)
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
	return storeTotal
}

func (productGroup ProductGroup) GetCategoryPrice(category string) (total float64) {
	for _, product := range productGroup {
		fmt.Println(category, "product:", product.Name)
		total = total + product.Price
	}
	fmt.Println(category, "subtotal:", ToCurrency(total))
	return
}
