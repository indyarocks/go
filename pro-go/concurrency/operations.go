package main

import (
	"fmt"
	"time"
)

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64)
	for category, productGroup := range data {
		fmt.Println("Invoking Goroutine--")
		//storeTotal += productGroup.GetCategoryPrice(category)
		go productGroup.GetCategoryPrice(category, channel)
		fmt.Println("Invoked Goroutine--")
	}
	for i := 0; i < len(data); i++ {
		fmt.Println(time.Now(), "Pending receive from channel----")
		storeTotal += <-channel
		fmt.Println(time.Now(), "Receive complete from channel----")
		time.Sleep(time.Second)
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}

func (productGroup ProductGroup) GetCategoryPrice(category string, channel chan float64) {
	var total float64
	for _, product := range productGroup {
		fmt.Println("Category:", category, "--product:", product.Name)
		total = total + product.Price
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println(time.Now(), category, "Sending to channel ----")
	fmt.Println(category, "subtotal:", ToCurrency(total))
	channel <- total
	fmt.Println(time.Now(), category, "Sent to channel ----")
}
