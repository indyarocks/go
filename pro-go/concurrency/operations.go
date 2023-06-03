package main

import (
	"fmt"
	"time"
)

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64, 2)
	for category, productGroup := range data {
		fmt.Println("Invoking Goroutine--")
		//storeTotal += productGroup.GetCategoryPrice(category)
		go productGroup.GetCategoryPrice(category, channel)
		fmt.Println("Invoked Goroutine--")
	}
	for i := 0; i < len(data); i++ {
		fmt.Println(len(channel), cap(channel))
		fmt.Println("----Channel read pending", len(channel), "items in buffer, size", cap(channel))
		storeTotal += <-channel
		fmt.Println("----Channel read complete", len(channel), "items in buffer, size", cap(channel))
		time.Sleep(time.Second)
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}

func (productGroup ProductGroup) GetCategoryPrice(category string, channel chan float64) {
	var total float64
	for _, product := range productGroup {
		//fmt.Println("Category:", category, "--product:", product.Name)
		total = total + product.Price
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println(category, "Channel sending", ToCurrency(total))
	channel <- total
	fmt.Println(category, "Channel send complete")
}
