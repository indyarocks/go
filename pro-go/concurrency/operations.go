package main

import (
	"fmt"
	"time"
)

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	for category, productGroup := range data {
		//storeTotal += productGroup.GetCategoryPrice(category)
		go productGroup.GetCategoryPrice(category)
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}

func (productGroup ProductGroup) GetCategoryPrice(category string) (total float64) {
	for _, product := range productGroup {
		fmt.Println(category, "product:", product.Name)
		total = total + product.Price
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println(category, "subtotal:", ToCurrency(total))
	return
}
