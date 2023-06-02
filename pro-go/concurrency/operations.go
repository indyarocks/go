package main

import (
	"fmt"
)

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64)
	for category, productGroup := range data {
		//storeTotal += productGroup.GetCategoryPrice(category)
		go productGroup.GetCategoryPrice(category, channel)
	}
	for i := 0; i < len(data); i++ {
		storeTotal += <-channel
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}

func (productGroup ProductGroup) GetCategoryPrice(category string, channel chan float64) {
	var total float64
	for _, product := range productGroup {
		fmt.Println(category, "product:", product.Name)
		total = total + product.Price
		//time.Sleep(time.Millisecond * 100)
	}
	fmt.Println(category, "subtotal:", ToCurrency(total))
	channel <- total
}
